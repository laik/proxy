package proxy

import (
	"crypto/tls"
	"net/url"
	"time"
)

type Storage interface {
	Set(key, value string) error
	Get(key string) (string, error)
	Del(key string) error
	All() (map[string]string, error)
}

type serverOption struct {
	addr string
	stge Storage
}

type ServerOptions func(*serverOption)

// StorageServerOption is duration storage
func AddrServerOption(addr string) ServerOptions {
	return func(s *serverOption) {
		s.addr = addr
	}
}

// StorageServerOption is duration storage
func StorageServerOption(stge Storage) ServerOptions {
	return func(s *serverOption) {
		s.stge = stge
	}
}

type listenerOption struct {
	Addr string
}

type ListenerOptions func(*listenerOption)

func AddrListenerOption(addr string) ListenerOptions {
	return func(l *listenerOption) {
		l.Addr = addr
	}
}

type handlerOption struct{}

type HandlerOptions func(*handlerOption)

// DialOptions describes the options for Transporter.Dial.
type DialOptions struct {
	Timeout time.Duration
	// Chain   *Chain
}

// DialOption allows a common way to set DialOptions.
type DialOption func(opts *DialOptions)

// TimeoutDialOption specifies the timeout used by Transporter.Dial
func TimeoutDialOption(timeout time.Duration) DialOption {
	return func(opts *DialOptions) {
		opts.Timeout = timeout
	}
}

// HandshakeOptions describes the options for handshake.
type HandshakeOptions struct {
	Addr       string
	Host       string
	User       *url.Userinfo
	Timeout    time.Duration
	Interval   time.Duration
	Retry      int
	TLSConfig  *tls.Config
	KCPConfig  *KCPConfig
	QUICConfig *QUICConfig
}

// HandshakeOption allows a common way to set HandshakeOptions.
type HandshakeOption func(opts *HandshakeOptions)

// AddrHandshakeOption specifies the server address
func AddrHandshakeOption(addr string) HandshakeOption {
	return func(opts *HandshakeOptions) {
		opts.Addr = addr
	}
}

// HostHandshakeOption specifies the hostname
func HostHandshakeOption(host string) HandshakeOption {
	return func(opts *HandshakeOptions) {
		opts.Host = host
	}
}

// UserHandshakeOption specifies the user used by Transporter.Handshake
func UserHandshakeOption(user *url.Userinfo) HandshakeOption {
	return func(opts *HandshakeOptions) {
		opts.User = user
	}
}

// TimeoutHandshakeOption specifies the timeout used by Transporter.Handshake
func TimeoutHandshakeOption(timeout time.Duration) HandshakeOption {
	return func(opts *HandshakeOptions) {
		opts.Timeout = timeout
	}
}

// IntervalHandshakeOption specifies the interval time used by Transporter.Handshake
func IntervalHandshakeOption(interval time.Duration) HandshakeOption {
	return func(opts *HandshakeOptions) {
		opts.Interval = interval
	}
}

// RetryHandshakeOption specifies the times of retry used by Transporter.Handshake
func RetryHandshakeOption(retry int) HandshakeOption {
	return func(opts *HandshakeOptions) {
		opts.Retry = retry
	}
}

// TLSConfigHandshakeOption specifies the TLS config used by Transporter.Handshake
func TLSConfigHandshakeOption(config *tls.Config) HandshakeOption {
	return func(opts *HandshakeOptions) {
		opts.TLSConfig = config
	}
}

// KCPConfigHandshakeOption specifies the KCP config used by KCP handshake
func KCPConfigHandshakeOption(config *KCPConfig) HandshakeOption {
	return func(opts *HandshakeOptions) {
		opts.KCPConfig = config
	}
}

// QUICConfigHandshakeOption specifies the QUIC config used by QUIC handshake
func QUICConfigHandshakeOption(config *QUICConfig) HandshakeOption {
	return func(opts *HandshakeOptions) {
		opts.QUICConfig = config
	}
}

// ConnectOptions describes the options for Connector.Connect.
type ConnectOptions struct {
	Addr    string
	Timeout time.Duration
	User    *url.Userinfo
}

// ConnectOption allows a common way to set ConnectOptions.
type ConnectOption func(opts *ConnectOptions)

// AddrConnectOption specifies the corresponding address of the target.
func AddrConnectOption(addr string) ConnectOption {
	return func(opts *ConnectOptions) {
		opts.Addr = addr
	}
}

// TimeoutConnectOption specifies the timeout for connecting to target.
func TimeoutConnectOption(timeout time.Duration) ConnectOption {
	return func(opts *ConnectOptions) {
		opts.Timeout = timeout
	}
}

// UserConnectOption specifies the user info for authentication.
func UserConnectOption(user *url.Userinfo) ConnectOption {
	return func(opts *ConnectOptions) {
		opts.User = user
	}
}
