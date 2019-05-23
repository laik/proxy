package proxy

import "net"

// Client is a proxy client.
// A client is divided into two layers: connector and transporter.
// Connector is responsible for connecting to the destination address through this proxy.
// Transporter performs a handshake with this proxy.
type Client struct {
	Connector   Connector
	Transporter Transporter
}

// Dial connects to the target address.
func (c *Client) Dial(addr string, options ...DialOption) (net.Conn, error) {
	return c.Transporter.Dial(addr, options...)
}

// Handshake performs a handshake with the proxy over connection conn.
func (c *Client) Handshake(conn net.Conn, options ...HandshakeOption) (net.Conn, error) {
	return c.Transporter.Handshake(conn, options...)
}

// Connect connects to the address addr via the proxy over connection conn.
func (c *Client) Connect(conn net.Conn, addr string, options ...ConnectOption) (net.Conn, error) {
	return c.Connector.Connect(conn, addr, options...)
}

// Connector is responsible for connecting to the destination address.
type Connector interface {
	Connect(conn net.Conn, addr string, options ...ConnectOption) (net.Conn, error)
}

// Transporter is responsible for handshaking with the proxy server.
type Transporter interface {
	Dial(addr string, options ...DialOption) (net.Conn, error)
	Handshake(conn net.Conn, options ...HandshakeOption) (net.Conn, error)
	// Indicate that the Transporter supports multiplex
	Multiplex() bool
}
