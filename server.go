package main

import (
	"log"
	"net/http"

	"golang.org/x/net/websocket"
)

type server struct {
	mb MessageBoard
}

func RunServer() {
	s := server{}

	http.HandleFunc("/", s.messageBoardHandler)

	http.HandleFunc("/style.css", s.cssHandler)

	http.HandleFunc("/script.js", s.jsHandler)

	http.HandleFunc("/message", s.messageHandler)

	http.Handle("/echo", websocket.Handler(EchoServer))

	log.Fatal(http.ListenAndServe(":8080", nil))
}
