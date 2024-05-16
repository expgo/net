// Copyright (C) 2019 The Syncthing Authors.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this file,
// You can obtain one at https://mozilla.org/MPL/2.0/.

//go:build go1.15 && !noquic
// +build go1.15,!noquic

package connections

import (
	"context"
	"crypto/tls"
	"fmt"
	"net"
	"net/url"
	"time"

	"github.com/quic-go/quic-go"

	"github.com/expgo/net/config"
	"github.com/expgo/net/connections/registry"
	"github.com/expgo/net/protocol"
)

const (
	quicPriority = 100

	// The timeout for connecting, accepting and creating the various
	// streams.
	quicOperationTimeout = 10 * time.Second
)

func init() {
	factory := &quicDialerFactory{}
	for _, scheme := range []string{"quic", "quic4", "quic6"} {
		dialers[scheme] = factory
	}
}

type quicDialer struct {
	commonDialer
	registry *registry.Registry
}

func (d *quicDialer) Dial(ctx context.Context, _ protocol.DeviceID, uri *url.URL) (internalConn, error) {
	uri = fixupPort(uri, config.DefaultQUICPort)

	network := quicNetwork(uri)

	addr, err := net.ResolveUDPAddr(network, uri.Host)
	if err != nil {
		return internalConn{}, err
	}

	var conn net.PacketConn
	// We need to track who created the conn.
	// Given we always pass the connection to quic, it assumes it's a remote connection it never closes it,
	// So our wrapper around it needs to close it, but it only needs to close it if it's not the listening connection.
	var createdConn net.PacketConn
	listenConn := d.registry.Get(uri.Scheme, packetConnUnspecified)
	if listenConn != nil {
		conn = listenConn.(net.PacketConn)
	} else {
		if packetConn, err := net.ListenPacket("udp", ":0"); err != nil {
			return internalConn{}, err
		} else {
			conn = packetConn
			createdConn = packetConn
		}
	}

	ctx, cancel := context.WithTimeout(ctx, quicOperationTimeout)
	defer cancel()

	session, err := quic.Dial(ctx, conn, addr, d.tlsCfg, quicConfig)
	if err != nil {
		if createdConn != nil {
			_ = createdConn.Close()
		}
		return internalConn{}, fmt.Errorf("dial: %w", err)
	}

	stream, err := session.OpenStreamSync(ctx)
	if err != nil {
		// It's ok to close these, this does not close the underlying packetConn.
		_ = session.CloseWithError(1, err.Error())
		if createdConn != nil {
			_ = createdConn.Close()
		}
		return internalConn{}, fmt.Errorf("open stream: %w", err)
	}

	return newInternalConn(&quicTlsConn{session, stream, createdConn}, connTypeQUICClient, quicPriority), nil
}

type quicDialerFactory struct{}

func (quicDialerFactory) New(opts config.OptionsConfiguration, tlsCfg *tls.Config, registry *registry.Registry) genericDialer {
	// So the idea is that we should probably try dialing every 20 seconds.
	// However it would still be nice if this was adjustable/proportional to ReconnectIntervalS
	// But prevent something silly like 1/3 = 0 etc.
	quicInterval := opts.ReconnectIntervalS / 3
	if quicInterval < 10 {
		quicInterval = 10
	}
	return &quicDialer{
		commonDialer: commonDialer{
			reconnectInterval: time.Duration(quicInterval) * time.Second,
			tlsCfg:            tlsCfg,
		},
		registry: registry,
	}
}

func (quicDialerFactory) Priority() int {
	return quicPriority
}

func (quicDialerFactory) AlwaysWAN() bool {
	return false
}

func (quicDialerFactory) Valid(_ config.Configuration) error {
	// Always valid
	return nil
}

func (quicDialerFactory) String() string {
	return "QUIC Dialer"
}
