package connection

import (
	"fmt"
	"net"
)

type Connection struct {
	conn net.Conn
	isConnected bool
} 

func Connect(ip, port string) (*Connection, error) {
	sock, err := net.Dial("tcp", ip+":"+port)
	
	if err != nil {
		return nil, err
	}

	conn := &Connection{
		conn: sock,
		isConnected: true,	
	}
	
	return conn, nil
}

func createConnection(conn net.Conn) (*Connection) {
	return &Connection{
		conn: conn,
		isConnected: true,
	}
}

func (conn *Connection) IsConnected() bool {
	return conn.isConnected;
}

func (conn *Connection) Disconnect() error {
	conn.isConnected = false
	fmt.Println("Disconnected")
	return conn.conn.Close()
}

func (conn *Connection) Receive(bytes_count uint32) ([]byte, error) {
	readBuffer := make([]byte, bytes_count)
	var cursor uint32 = 0

	for cursor < bytes_count {
		n, err := conn.conn.Read(readBuffer[cursor:])
		
		if err != nil {
			return nil, fmt.Errorf("Error during read")
		}
		
		cursor += uint32(n)
	}

	return readBuffer, nil
}

func (conn *Connection) Send(bytes []byte) error {

	for len(bytes) > 0 {
		n, err := conn.conn.Write(bytes)

		if err != nil {
			return fmt.Errorf("Error during sending")
		}
		
		bytes = bytes[n:]
	}
	return nil
}