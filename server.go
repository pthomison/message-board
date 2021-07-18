package main

import (
	"log"
	"net/http"
)

func RunServer() {
	http.HandleFunc("/", messageBoardHandler)

	http.HandleFunc("/style.css", cssHandler)

	http.HandleFunc("/message", messageHandler)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
