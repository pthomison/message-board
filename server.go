package main

import (
	"fmt"
	"html"
	"log"
	"net/http"
)

type Employee struct {
	FirstName   string
	LastName    string
	TotalLeaves int
	LeavesTaken int
}

func RunServer() {
	http.HandleFunc("/", messageBoardHandler)

	http.HandleFunc("/style.css", cssHandler)

	http.HandleFunc("/bar", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}
