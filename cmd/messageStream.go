package cmd

import (
	"fmt"

	"golang.org/x/net/websocket"
)

// Echo the data received on the WebSocket.
func (s *server) MessageStream(ws *websocket.Conn) {
	fmt.Printf("jsonServer %#v\n", ws.Config())
	for k, v := range s.wb {
		var msg T
		err := websocket.JSON.Receive(ws, &msg)
		if err != nil {
			fmt.Println(err)
			break
		}
		fmt.Printf("recv:%#v\n", msg)
		err = websocket.JSON.Send(ws, msg)
		if err != nil {
			fmt.Println(err)
			break
		}
		fmt.Printf("send:%#v\n", msg)
	}
	fmt.Println("jsonServer finished")
}

func Json(ws *websocket.Conn) {

}
