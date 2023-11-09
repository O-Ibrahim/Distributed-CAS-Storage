package p2p

//Peer is an interface that represents the remote node
type Peer interface {
}

//Transport handles the communication between the nodes in the network
type Transport interface {
	Listen() error
}
