// Copyright (C) 2019 The Syncthing Authors.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this file,
// You can obtain one at http://mozilla.org/MPL/2.0/.

//go:build go1.15 && !noquic
// +build go1.15,!noquic

package connections

import (
	"crypto/tls"
	"net"
	"net/url"
	"time"

	"github.com/quic-go/quic-go"
)

var (
	quicConfig = &quic.Config{
		// TODO py
		//ConnectionIDLength: 4,
		MaxIdleTimeout:  30 * time.Second,
		KeepAlivePeriod: 15 * time.Second,
	}
)

func quicNetwork(uri *url.URL) string {
	switch uri.Scheme {
	case "quic4":
		return "udp4"
	case "quic6":
		return "udp6"
	default:
		return "udp"
	}
}

type quicTlsConn struct {
	quic.Connection
	quic.Stream
	// If we created this connection, we should be the ones closing it.
	createdConn net.PacketConn
}

func (q *quicTlsConn) Close() error {
	sterr := q.Stream.Close()
	seerr := q.Connection.CloseWithError(0, "closing")
	var pcerr error
	if q.createdConn != nil {
		pcerr = q.createdConn.Close()
	}
	if sterr != nil {
		return sterr
	}
	if seerr != nil {
		return seerr
	}
	return pcerr
}

func (q *quicTlsConn) ConnectionState() tls.ConnectionState {
	return q.Connection.ConnectionState().TLS
}

func packetConnUnspecified(conn interface{}) bool {
	// Since QUIC connections are wrapped, we can't do a simple typecheck
	// on *net.UDPAddr here.
	addr := conn.(net.PacketConn).LocalAddr()
	host, _, err := net.SplitHostPort(addr.String())
	return err == nil && net.ParseIP(host).IsUnspecified()
}
