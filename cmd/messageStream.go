package cmd

import (
	"fmt"

	utils "github.com/pthomison/golang-utils"
	"golang.org/x/net/websocket"
)

func (s *server) MessageStream(ws *websocket.Conn) {
	fmt.Printf("MessageStream %#v\n", ws.Config())

	msgChannel := make(chan Message)
	s.mb.Connections[ws] = msgChannel
	defer delete(s.mb.Connections, ws)

	for _, v := range s.mb.Messages {
		SendMessage(ws, v)
	}

	for m := range msgChannel {
		SendMessage(ws, m)
	}

	fmt.Println("MessageStream finished")
}

func SendMessage(ws *websocket.Conn, m Message) {
	err := websocket.JSON.Send(ws, m)
	utils.Check(err)
}

func (s *server) MessageReceive(ws *websocket.Conn) {
	fmt.Printf("MessageReceive %#v\n", ws.Config())

	var msg Message
	websocket.JSON.Receive(ws, &msg)

	s.mb.Messages = append(s.mb.Messages, msg)

	for _, ch := range s.mb.Connections {
		ch <- msg
	}

	fmt.Printf("Message Recieved: %+v\n", msg)

	fmt.Println("MessageReceive finished")
}
