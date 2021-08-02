package cmd

import (
	"fmt"

	utils "github.com/pthomison/golang-utils"
	"golang.org/x/net/websocket"
)

func (s *server) MessageStream(ws *websocket.Conn) {
	fmt.Printf("MessageStream %#v\n", ws.Config())

	for _, v := range s.mb.Messages {
		SendMessage(ws, v)
	}

	for m := range s.mc {
		SendMessage(ws, m)
	}

	fmt.Println("MessageStream finished")
}

func SendMessage(ws *websocket.Conn, m Message) {
	err := websocket.JSON.Send(ws, m)
	utils.Check(err)
}
