package p2p

import "errors"

var (
	//ErrInvalidHandShake is returned if the handshake between the system and the target does not pass
	ErrInvalidHandShake = errors.New("invalid handshake")
)

// HandShakeFunc is a function that runs in order to confirm the connection with peer
type HandShakeFunc func(Peer) error

// NOPHandShakeFunc is a no operation func that returns nil
func NOPHandShakeFunc(p Peer) error {
	return nil
}
