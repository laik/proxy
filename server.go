package proxy

import (
	"net"
	"sync"
)

var (
	tinyBufferSize   = 128
	smallBufferSize  = 1 * 1024  // 1KB small buffer
	mediumBufferSize = 8 * 1024  // 8KB medium buffer
	largeBufferSize  = 32 * 1024 // 32KB large buffer
)

var (
	sPool = sync.Pool{
		New: func() interface{} {
			return make([]byte, smallBufferSize)
		},
	}
	mPool = sync.Pool{
		New: func() interface{} {
			return make([]byte, mediumBufferSize)
		},
	}
	lPool = sync.Pool{
		New: func() interface{} {
			return make([]byte, largeBufferSize)
		},
	}
)

// Listener is a proxy server listener, just like a net.Listener.
type Listener interface {
	net.Listener
}

type tcpListener struct {
	net.Listener
}

// Accepter represents a network endpoint that can accept connection from peer.
type Accepter interface {
	Accept() (net.Conn, error)
}

// Handler is a proxy server handler
type Handler interface {
	Init(options ...HandlerOptions)
	Handle(net.Conn)
}

type ServerMux map[Listener]Handler

// Server is a proxy server.
type Server struct {
	mu      sync.Mutex
	muxs    map[string]ServerMux
	options *serverOption
}

// Init intializes server with given options.
func (s *Server) Init(opts ...ServerOptions) {
	if s.options == nil {
		s.options = &serverOption{}
	}
	for _, opt := range opts {
		opt(s.options)
	}
	
}

func (s *Server) add(l Listener, listenerOpts []ListenerOptions, h Handler, handlerOpts []HandlerOptions) error {

	return nil
}

func (s *Server) remove() error {
	return nil
}

func (s *Server) Serve(opts ...ServerOptions) error {
	s.Init(opts...)

	return nil
}
