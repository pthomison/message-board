package cmd

import (
	"embed"
	"log"
	"net/http"

	"golang.org/x/net/websocket"
)

type server struct {
	mb        MessageBoard
	uiAssests embed.FS
	mc        chan Message
}

func RunServer(uiAssests embed.FS) {
	s := server{
		uiAssests: uiAssests,
		mb:        bones(),
		mc:        make(chan Message, 100),
	}

	http.HandleFunc("/", s.messageBoardHandler)

	staticFS := http.FS(s.uiAssests)
	httpFS := http.FileServer(staticFS)
	http.Handle("/static/", http.StripPrefix("/static/", httpFS))

	http.HandleFunc("/message", s.messageHandler)

	http.Handle("/message-stream", websocket.Handler(s.MessageStream))

	http.Handle("/message-receive", websocket.Handler(s.MessageReceive))

	log.Fatal(http.ListenAndServe(":8080", nil))
}
