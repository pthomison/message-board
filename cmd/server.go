package cmd

import (
	"io/fs"
	"log"
	"net/http"

	"golang.org/x/net/websocket"
)

type server struct {
	mb        MessageBoard
	uiAssests fs.FS
}

func RunServer(uiAssests fs.FS) {
	s := server{
		uiAssests: uiAssests,
		mb:        bones(),
	}

	http.HandleFunc("/", s.messageBoardHandler)

	http.HandleFunc("/style.css", s.cssHandler)

	http.HandleFunc("/script.js", s.jsHandler)

	http.HandleFunc("/message", s.messageHandler)

	http.Handle("/echo", websocket.Handler(EchoServer))

	log.Fatal(http.ListenAndServe(":8080", nil))
}
