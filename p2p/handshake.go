package p2p

//HandShakeFunc is a function that runs in order to confirm the connection with peer
type HandShakeFunc func(any) error

//NOPHandShakeFunc is a no operation func that returns nil
func NOPHandShakeFunc(a any) error {
	return nil
}
