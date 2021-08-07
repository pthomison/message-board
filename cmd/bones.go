package cmd

import (
	"time"

	"golang.org/x/net/websocket"
)

func bones() MessageBoard {

	testMsgs := []Message{
		Message{
			Author:    "John Doe",
			Message:   "Hello World This Is A Test",
			Date:      "12-12-12",
			Timestamp: time.Now(),
		},
		Message{
			Author:    "John Doe",
			Message:   "Hello World This Is B Test",
			Date:      "12-12-12",
			Timestamp: time.Now(),
		},
		Message{
			Author:    "John Doe",
			Message:   "Hello World This Is C Test",
			Date:      "12-12-12",
			Timestamp: time.Now(),
		},
	}

	var messages []Message
	var connections map[*websocket.Conn]chan Message

	for i := 0; i < 1; i++ {
		tmp := make([]Message, len(testMsgs))
		copy(tmp, testMsgs)
		messages = append(messages, tmp...)
	}

	connections = make(map[*websocket.Conn]chan Message)

	return MessageBoard{
		Messages:    messages,
		Connections: connections,
	}
}
