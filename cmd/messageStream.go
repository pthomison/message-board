package cmd

import (
	"fmt"

	"golang.org/x/net/websocket"
)

// Echo the data received on the WebSocket.
func (s *server) MessageStream(ws *websocket.Conn) {
	fmt.Printf("jsonServer %#v\n", ws.Config())
	for _, v := range s.mb.Messages {
		// var msg Message
		err := websocket.JSON.Send(ws, v)
		fmt.Println("error", err)
		if err != nil {
			fmt.Println("error", err)
			break
		}
		fmt.Printf("send:%#v\n", v)
	}
	fmt.Println("jsonServer finished")
}

func Json(ws *websocket.Conn) {

}
