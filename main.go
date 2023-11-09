package main

import (
	"log"

	"github.com/O-Ibrahim/Distributed-CAS-Storage/p2p"
)

func main() {
	transport := p2p.NewTCPTransport(":3000")
	if err := transport.Listen(); err != nil {
		log.Fatal(err)
	}
	select {}
}
