package connection

import (
	"chat-tcp/assert"
	"encoding/binary"
	"errors"
)

const HEADER_SIZE = 4
const MAX_MESSAGE_SIZE = 2097152

type Framer struct {
	tcp *Connection
}

type header struct {
	size uint32 
}

func CreateFramer(tcp *Connection) *Framer {
	return &Framer {
		tcp: tcp,
	}
}

func unmarshallHeader(bytes []byte) header {
	assert.Assert(len(bytes) >= HEADER_SIZE, "Byte slice to small to get header");

	size := binary.BigEndian.Uint32(bytes)	
	return header {
		size: size,
	}
}

func marshallHeader(header header) []byte {
	bytes := make([]byte, HEADER_SIZE)
	binary.BigEndian.PutUint32(bytes[0:], header.size)
	return bytes
}

// Jeśli ta funkcja zwróci error - połączenie powinno zostać zerwane
// TODO: wykorzystanie interfejsu reader do łatwiejszego testowania
func (self *Framer) Receive() ([]byte, error) {
	bytes, err := self.tcp.Receive(HEADER_SIZE) 
	if err != nil {
		return nil, err
	}

	header := unmarshallHeader(bytes)

	if header.size > MAX_MESSAGE_SIZE {
		return nil, errors.New("frame is too big")
	}

	bytesPending := header.size

	frame := make([]byte, 0)

	for bytesPending > 0 {
		bytesToRead := min(4096, bytesPending)
		bytes, err := self.tcp.Receive(bytesToRead)
		
		if err != nil {
			return nil, err
		}

		bytesPending -= bytesToRead

		frame = append(frame, bytes...)
	}

	return frame, nil
}

func (self *Framer) Send(bytes []byte) error {
	header := header {
		size: uint32(len(bytes)),
	}

	if header.size > MAX_MESSAGE_SIZE {
		return errors.New("Message is too big to send")
	}

	buffer := marshallHeader(header)

	buffer = append(buffer, bytes...)
	
	err := self.tcp.Send(buffer)

	if err != nil {
		return err
	}
	
	return nil
}