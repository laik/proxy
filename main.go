package main

import (
	"io"
	"net"
	"strings"
	"sync"
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
	bufpool = sync.Pool{
		New: func() interface{} {
			buf := make([]uint8, 65535)
			return &buf
		},
	}

	lazy = func() []uint8 {
		buf := bufpool.Get().(*[]uint8)
		return *buf
	}

	putbuf = func(p []uint8) {
		bufpool.Put(p)
	}
)

type (
	ModeType   uint8
	MediumType uint8
)

const (
	Server ModeType = iota
	Client

	KCP MediumType = iota
	QUIC
)

type option struct {
	Mode      ModeType   `json:"mode"`
	Listen    string     `json:"listen"`
	Medium    MediumType `json:"medium"`
	Target    string     `json:"target"`
	KeepAlive int        `json:"keepalive"`
}

type Option struct {
	f func(*option)
}

func WithMode(n ModeType) Option {
	return Option{func(o *option) {
		o.Mode = n
	}}
}

func WithMedium(m MediumType) Option {
	return Option{func(o *option) {
		o.Medium = m
	}}
}

type Logger interface {
	Info(string, ...interface{})
	Error(string, ...interface{})
}

// A transport.
var transport = func(dst, src io.ReadWriteCloser) {
	copy := func(dst io.Writer, src io.Reader) chan struct{} {
		die := make(chan struct{})
		go func() {
			if wt, ok := src.(io.WriterTo); ok {
				wt.WriteTo(dst)
				close(die)
			} else if rt, ok := dst.(io.ReaderFrom); ok {
				rt.ReadFrom(src)
				close(die)
			} else {
				buf := lazy()
				io.CopyBuffer(dst, src, buf)
				putbuf(buf)
				close(die)
			}
		}()
		return die
	}
	select {
	case <-copy(dst, src):
	case <-copy(src, dst):
	}
}

// Stream is the interface implemented by QUIC/KCP streams
type Stream interface {
	StreamID() uint64
	io.ReadWriteCloser
	SetDeadline(t time.Time) error
}

// A Session is a QUIC/KCP connection between two peers.
type Session interface {
	AcceptStream() (Stream, error)
	OpenStream() (Stream, error)
	OpenStreamSync() (Stream, error)
	LocalAddr() net.Addr
	RemoteAddr() net.Addr
	io.Closer
}

type TunnelListener interface {
	Accept() (Session, error)
	Close() error
	Addr() net.Addr
}

// A Listener for incoming QUIC/KCP connections
type Listeners struct {
	netListener    net.Listener
	tunnelListener TunnelListener
}

// A Hahdler handle stream.
type Handler interface {
	handle(p1, p2 io.ReadWriteCloser, logger Logger)
}

type tcpListener struct {
	net.Listener
}

// TCPListener creates a Listener for TCP proxy server.
func TCPListener(addr string) (net.Listener, error) {
	laddr, err := net.ResolveTCPAddr("tcp", addr)
	if err != nil {
		return nil, err
	}
	ln, err := net.ListenTCP("tcp", laddr)
	if err != nil {
		return nil, err
	}
	return &tcpListener{tcpKeepAliveListener{ln}}, nil
}

type tcpKeepAliveListener struct {
	*net.TCPListener
}

func (ln tcpKeepAliveListener) Accept() (c net.Conn, err error) {
	tc, err := ln.AcceptTCP()
	if err != nil {
		return
	}
	tc.SetKeepAlive(true)
	tc.SetKeepAlivePeriod(KeepAliveTime)
	return tc, nil
}

type Dialer interface {
	Dial(string, ...Option)
}

// A Proxy for transport local steam to define route ends.
type Proxy struct {
	listener *Listeners
	dialer   Dialer
	opt      *option
}

func NewProxy(src, dst string, opts ...Option) (proxy *Proxy, err error) {
	o := &option{}
	host, port, err := net.SplitHostPort(src)
	if err != nil {
		return nil, err
	}
	o.Listen = strings.Join([]string{host, port}, ":")

	host, port, err = net.SplitHostPort(dst)
	if err != nil {
		return nil, err
	}
	o.Target = strings.Join([]string{host, port}, ":")

	for _, opt := range opts {
		opt.f(o)
	}

	var b builder = &proxyBuilder{}

	switch o.Mode {
	case Server:
		proxy, err = b.withOption(o).buildServer()
		if err != nil {
			return nil, err
		}
		return proxy, nil
	case Client:
		proxy, err = b.withOption(o).buildClient()
		if err != nil {
			return nil, err
		}
		return proxy, nil
	}

	return nil, nil
}

func main() {

}
