package p2p

import "net"

//Message represents data being sent over the trasport
type Message struct {
	From net.Addr
	Payload []byte
}