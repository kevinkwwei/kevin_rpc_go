package rpc_transport

import "net"

const DefaultPayloadLength = 1024
const MaxPayloadLength = 4 * 1024 * 1024

// 网络连接中的完整帧
type Framer interface {
	ReadFrame(net.Conn) ([]byte, error)
}

type framer struct {
	buffer  []byte
	counter int
}

func (f framer) ReadFrame(conn net.Conn) ([]byte, error) {
	// decode by codec and return the result
	panic("to do here")
}

func NewFramer() Framer {
	return &framer{
		buffer: make([]byte, DefaultPayloadLength),
	}
}

func (f *framer) Resize() {
	f.buffer = make([]byte, len(f.buffer)*2)
}
