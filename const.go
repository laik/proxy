package proxy

import (
	"crypto/tls"
	"time"
)

var (
	// KeepAliveTime is the keep alive time period for TCP connection.
	KeepAliveTime = 180 * time.Second
	// DialTimeout is the timeout of dial.
	DialTimeout = 5 * time.Second
	// HandshakeTimeout is the timeout of handshake.
	HandshakeTimeout = 5 * time.Second
	// ConnectTimeout is the timeout for connect.
	ConnectTimeout = 5 * time.Second
	// ReadTimeout is the timeout for reading.
	ReadTimeout = 10 * time.Second
	// WriteTimeout is the timeout for writing.
	WriteTimeout = 10 * time.Second
	// PingTimeout is the timeout for pinging.
	PingTimeout = 30 * time.Second
	// PingRetries is the reties of ping.
	PingRetries = 1
	// default udp node TTL in second for udp port forwarding.
	defaultTTL = 60 * time.Second
)

var (
	// DefaultTLSConfig is a default TLS config for internal use.
	DefaultTLSConfig *tls.Config
)
