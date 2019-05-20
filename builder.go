package main

import "github.com/pkg/errors"

type builder interface {
	buildServer() (*Proxy, error)
	buildClient() (*Proxy, error)
	withOption(*option) builder
}

type proxyBuilder struct {
	o *option
}

func (b *proxyBuilder) buildServer() (*Proxy, error) {
	tcpLn, err := TCPListener(b.o.Listen)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	proxy := &Proxy{}
	proxy.listener.netListener = tcpLn
	return proxy, nil
}

func (b *proxyBuilder) buildTcpListener() error {
	return nil
}

func (b *proxyBuilder) buildClient() (*Proxy, error) {
	return nil, nil
}

func (b *proxyBuilder) withOption(opt *option) builder {
	b.o = opt
	return b
}
