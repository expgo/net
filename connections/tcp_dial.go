// Copyright (C) 2016 The Syncthing Authors.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this file,
// You can obtain one at https://mozilla.org/MPL/2.0/.

package connections

import (
	"context"
	"crypto/tls"
	"net/url"
	"time"

	"github.com/expgo/net/config"
	"github.com/expgo/net/connections/registry"
	"github.com/expgo/net/dialer"
	"github.com/expgo/net/protocol"
)

const tcpPriority = 10

func init() {
	factory := &tcpDialerFactory{}
	for _, scheme := range []string{"tcp", "tcp4", "tcp6"} {
		dialers[scheme] = factory
	}
}

type tcpDialer struct {
	commonDialer
	registry *registry.Registry
}

func (d *tcpDialer) Dial(ctx context.Context, _ protocol.DeviceID, uri *url.URL) (internalConn, error) {
	uri = fixupPort(uri, config.DefaultTCPPort)

	timeoutCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()
	conn, err := dialer.DialContextReusePortFunc(d.registry)(timeoutCtx, uri.Scheme, uri.Host)
	if err != nil {
		return internalConn{}, err
	}

	err = dialer.SetTCPOptions(conn)
	if err != nil {
		l.Debugln("Dial (BEP/tcp): setting tcp options:", err)
	}

	err = dialer.SetTrafficClass(conn, d.trafficClass)
	if err != nil {
		l.Debugln("Dial (BEP/tcp): setting traffic class:", err)
	}

	tc := tls.Client(conn, d.tlsCfg)
	err = tlsTimedHandshake(tc)
	if err != nil {
		tc.Close()
		return internalConn{}, err
	}

	return newInternalConn(tc, connTypeTCPClient, tcpPriority), nil
}

type tcpDialerFactory struct{}

func (tcpDialerFactory) New(opts config.OptionsConfiguration, tlsCfg *tls.Config, registry *registry.Registry) genericDialer {
	return &tcpDialer{
		commonDialer: commonDialer{
			trafficClass:      opts.TrafficClass,
			reconnectInterval: time.Duration(opts.ReconnectIntervalS) * time.Second,
			tlsCfg:            tlsCfg,
		},
		registry: registry,
	}
}

func (tcpDialerFactory) Priority() int {
	return tcpPriority
}

func (tcpDialerFactory) AlwaysWAN() bool {
	return false
}

func (tcpDialerFactory) Valid(_ config.Configuration) error {
	// Always valid
	return nil
}

func (tcpDialerFactory) String() string {
	return "TCP Dialer"
}
