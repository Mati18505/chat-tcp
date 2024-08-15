package main

import (
	"chat-tcp/connection"
	"fmt"
)

type serverData struct {
	server *connection.Server
	broadcast chan connection.Message
}

func main() {
	server, err := connection.CreateServer("localhost", "65000")
	if err != nil {
		fmt.Println("Cannot create server: ", err)
		return
	}

	serverData := serverData{
		server: server,
		broadcast: make(chan connection.Message),
	}

	go ProcessMessages(&serverData)

	server.Accept(func (client connection.Client) {
		framer := connection.CreateFramer(client.Connection)
		messanger := connection.CreateMessanger(framer) 
		defer server.Disconnect(client.Uuid) 

		for client.Connection.IsConnected() {
			message, err := messanger.Receive()
			if err != nil {
				fmt.Println(err)
				return
			}

			fmt.Println("received: ", message)	
			serverData.broadcast <- *message
		}
	})		
}

func ProcessMessages(server *serverData) {
	var message connection.Message

	for {
		select {
		case message = <-server.broadcast:
			err := server.server.Broadcast(message)

			if err != nil {
				fmt.Println(err)
				return
			}
		}
	}
}