package peerstream_multiplex

import (
	"errors"
	"net"

	smux "github.com/jbenet/go-peerstream/Godeps/_workspace/src/github.com/jbenet/go-stream-muxer"
	mp "github.com/jbenet/go-peerstream/Godeps/_workspace/src/github.com/whyrusleeping/go-multiplex"
)

var ErrUseServe = // Conn is a connection to a remote peer.
errors.New("not implemented, use Serve")

type conn struct {
	*mp.Multiplex
}

func (c *conn) Close() error {
	return c.Multiplex.Close()
}

func (c *conn) IsClosed() bool {
	return c.Multiplex.IsClosed()
}

// OpenStream creates a new stream.
func (c *conn) OpenStream() (smux.Stream, error) {
	return c.Multiplex.NewStream(), nil
}

// AcceptStream accepts a stream opened by the other side.
func (c *conn) AcceptStream() (smux.Stream, error) {
	return nil, ErrUseServe
}

// Serve starts listening for incoming requests and handles them
// using given StreamHandler
func (c *conn) Serve(handler smux.StreamHandler) {
	c.Multiplex.Serve(func(s *mp.Stream) {
		handler(s)
	})
}

// Transport is a go-peerstream transport that constructs
// multiplex-backed connections.
type Transport struct{}

// DefaultTransport has default settings for multiplex
var DefaultTransport = &Transport{}

func (t *Transport) NewConn(nc net.Conn, isServer bool) (smux.Conn, error) {
	return &conn{mp.NewMultiplex(nc, isServer)}, nil
}