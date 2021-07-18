package main

import (
	"fmt"
	"html"
	"log"
	"net/http"
)

// //go:embed ui/*.css
// var css embed.FS

func RunServer() {
	http.HandleFunc("/", messageBoardHandler)

	http.HandleFunc("/style.css", cssHandler)

	// http.Handle("/css/", http.FileServer(http.FS(css)))

	http.HandleFunc("/bar", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}
