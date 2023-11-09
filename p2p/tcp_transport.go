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

type TCPTransport struct {
	listenAddr    string
	listener      net.Listener
	shakeHands HandShakeFunc
	decoder Decoder

	mu    sync.RWMutex
	peers map[net.Addr]Peer
}

func NewTCPTransport(listenAddr string) *TCPTransport {
	return &TCPTransport{
		shakeHands: NOPHandShakeFunc,
		listenAddr:    listenAddr,
	}
}


func (t *TCPTransport) Listen() error {
	ln, err := net.Listen("tcp", t.listenAddr)
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

func (t *TCPTransport) handleConn(conn net.Conn) {
	if err := t.shakeHands(conn); err != nil {
		conn.Close()
	}

	peer := NewTCPPeer(conn, true)
	fmt.Printf("New incoming connection %+v\n", peer)

	for {
		 
	}

}
