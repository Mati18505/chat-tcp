package connection

type Messanger struct {
	framer *Framer
}

type Message struct {
	Message string
}

func CreateMessanger(framer *Framer) *Messanger {
	return &Messanger{
		framer: framer,
	}
}

// Jeśli ta funkcja zwróci error - połączenie powinno zostać zerwane
func (self *Messanger) Receive() (*Message, error) {
	bytes, err := self.framer.Receive()

	if err != nil {
		return nil, err
	}

	message := unmarshall(bytes)

	return message, nil
}

// Jeśli ta funkcja zwróci error - połączenie powinno zostać zerwane
func (self *Messanger) Send(message Message) error {
	bytes := marshall(message)
	return self.framer.Send(bytes)
}

func marshall(message Message) []byte {
	return []byte(message.Message)
}

func unmarshall(bytes []byte) (*Message) {
	return &Message{
		Message: string(bytes),
	}
}