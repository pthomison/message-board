package cmd

import (
	"golang.org/x/net/websocket"
	"fmt"
	"time"
)

func PingServer(ws *websocket.Conn) {
	fmt.Printf("PingServer %#v\n", ws.Config())
	for {
		msg := &Message{
			Author: "ted dancing",
			Message: "whats cooking good looking",
			Date:    "this morning",
			Timestamp: time.Now(),
		}

		err := websocket.JSON.Send(ws, msg)
		fmt.Println("error", err)
		if err != nil {
			fmt.Println("error", err)
			break
		}
		fmt.Printf("send:%#v\n", msg)

	    // Calling Sleep method
	    time.Sleep(2 * time.Second)

	}
	fmt.Println("PingServer finished")
}