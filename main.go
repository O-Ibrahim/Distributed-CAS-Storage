package main

import (
	"log"

	"github.com/O-Ibrahim/Distributed-CAS-Storage/p2p"
)

func main() {
	transport := p2p.NewTCPTransport(p2p.TCPTransportOptions{
		ListenAddr: ":3000",
		ShakeHands: p2p.NOPHandShakeFunc,
		Decoder:    p2p.NewDefaultDecoder(),
	})
	if err := transport.Listen(); err != nil {
		log.Fatal(err)
	}
	select {}
}
