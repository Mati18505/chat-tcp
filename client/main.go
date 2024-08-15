package main

import (
	"bufio"
	"chat-tcp/connection"
	"fmt"
	"os"
	"time"
)

func main() {

	const maxConnectionAttempt = 8
	var connectionAttempt = 0

	con, err := connection.Connect("localhost", "65000")

	for err != nil {
		time.Sleep(250 * time.Millisecond)
		con, err = connection.Connect("localhost", "65000")
		connectionAttempt++

		if connectionAttempt >= maxConnectionAttempt {
			fmt.Println("Connection timed out:", err)
			return
		}
	}

	framer := connection.CreateFramer(con)
	messanger := connection.CreateMessanger(framer)

	defer con.Disconnect()

	for con.IsConnected() {

		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		message := scanner.Text()

		err := messanger.Send(
		connection.Message{
			Message: message,
		})

		if err != nil {
			fmt.Println(err)
			return
		}
	}
}