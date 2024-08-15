package connection

import (
	"errors"
	"net"

	"github.com/google/uuid"
)

type Client struct {
	Uuid uuid.UUID
	Connection *Connection
}
type Server struct {
	listener net.Listener
	connections map[uuid.UUID]Client
}

func CreateServer(ip, port string) (*Server, error) {
	listener, err := net.Listen("tcp", ip+":"+port)

	if err != nil {
		return nil, err
	}

	return &Server{
		listener: listener,
		connections: make(map[uuid.UUID]Client),
	}, nil
}

// Accepts incoming client connections, on connection executes connectionHandler
func (self *Server) Accept(connectionHandler func(Client)) {
	for {
		conn, err := self.listener.Accept()
	
		if err != nil {
			continue;
		}
		connection := createConnection(conn)

		client, err := self.addClient(connection)
		if err != nil {
			continue 
		}

		go connectionHandler(client)
	}
}

func (self *Server) Broadcast(message Message) error {
	for uuid, client := range self.connections { 
		con := client.Connection
		framer := CreateFramer(con)
		messanger := CreateMessanger(framer) 
		err := messanger.Send(message)

		if err != nil {
			delete(self.connections, uuid) 
			con.Disconnect() 
			return err
		}
	}
	return nil
}

func (self *Server) Disconnect(uuid uuid.UUID) error {
	client, exist := self.connections[uuid]
	if !exist {
		return errors.New("Client with this UUID does not exist!")
	}

	client.Connection.Disconnect()
	delete(self.connections, uuid)

	return nil
}

func (self *Server) addClient(connection *Connection) (Client, error) {
	uuid, err := uuid.NewRandom()
	if err != nil {
		return Client{}, err
	}

	client := Client{
		Uuid: uuid,
		Connection: connection,
	}

	self.connections[uuid] = client
	return client, nil
}