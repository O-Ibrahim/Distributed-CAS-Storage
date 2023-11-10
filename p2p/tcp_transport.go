package p2p

import (
	"fmt"
	"net"
	"sync"
)

// TCPPeer represents the remote node over TCP connection
type TCPPeer struct {
	//the underlying connection of the peer
	conn net.Conn
	// if we dial an external connection = outbound = outbound true
	// if we accept an external connection = inbound = outbound false
	outBound bool
}

func NewTCPPeer(conn net.Conn, outBound bool) *TCPPeer {
	return &TCPPeer{
		conn:     conn,
		outBound: outBound,
	}
}

type TCPTransportOptions struct {
	ListenAddr string
	ShakeHands HandShakeFunc
	Decoder    Decoder
}

type TCPTransport struct {
	opts TCPTransportOptions

	listener net.Listener

	mu    sync.RWMutex
	peers map[net.Addr]Peer
}

func NewTCPTransport(opts TCPTransportOptions) *TCPTransport {
	return &TCPTransport{
		opts: opts,
	}
}

func (t *TCPTransport) Listen() error {
	ln, err := net.Listen("tcp", t.opts.ListenAddr)
	if err != nil {
		return err
	}

	t.listener = ln
	go t.startAcceptLoop()
	return nil
}

func (t *TCPTransport) startAcceptLoop() {
	for {
		conn, err := t.listener.Accept()
		if err != nil {
			fmt.Printf("tcp transport error accepting connection:  %s\n", err)
		}
		go t.handleConn(conn)
	}
}

type Temp struct {
}

func (t *TCPTransport) handleConn(conn net.Conn) {
	peer := NewTCPPeer(conn, true)
	if err := t.opts.ShakeHands(peer); err != nil {
		conn.Close()
		fmt.Printf("tcp error - handshake with client %s\n", conn.RemoteAddr())
		return
	}

	fmt.Printf("New incoming connection %+v\n", peer)
	msg := &Message{
		From: conn.RemoteAddr(),
	}
	for {
		if err := t.opts.Decoder.Decode(conn, msg); err != nil {
			fmt.Printf("tcp error - decode %s\n", err)
			conn.Close()
			continue
		}

		fmt.Printf("message %v\n", msg)

	}

}
