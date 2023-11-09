package p2p

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_TestTCPTransport(t *testing.T) {
	lsnAddr := ":4000"
	tr := NewTCPTransport(lsnAddr)
	assert.Equal(t, tr.listenAddr, lsnAddr)

	err := tr.Listen()
	assert.NoError(t, err)

}

func Test_GivenInvalidListenAddr_whenListening_throwError(t *testing.T) {
	lsnAddr := "4000"
	tr := NewTCPTransport(lsnAddr)
	assert.Equal(t, tr.listenAddr, lsnAddr)

	err := tr.Listen()
	assert.Error(t, err)
}
