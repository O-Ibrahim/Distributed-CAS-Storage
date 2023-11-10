package p2p

import (
	"encoding/gob"
	"io"
)

type Decoder interface {
	Decode(io.Reader, *Message) error
}

type GobDecoder struct{}

func NewGobDecoder() *GobDecoder {
	return &GobDecoder{}
}
func (decoder *GobDecoder) Decode(r io.Reader, msg *Message) error {
	return gob.NewDecoder(r).Decode(msg)
}

type DefaultDecoder struct {
}

func NewDefaultDecoder() *DefaultDecoder {
	return &DefaultDecoder{}
}

func (decoder *DefaultDecoder) Decode(r io.Reader, msg *Message) error {
	buf := make([]byte, 1024)
	n, err := r.Read(buf)
	if err != nil {
		return err
	}
	msg.Payload = buf[:n]
	return nil
}
