package main

import (
	"fmt"
	"net/http"
	"strings"
)

type Message struct {
	Author  string
	Message string
	Date    string
}

func messageHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		errorHandler(w, r, http.StatusNotFound)
	case http.MethodPost:
		messagePostHandler(w, r)
	case http.MethodPut:
		errorHandler(w, r, http.StatusNotFound)
	case http.MethodDelete:
		errorHandler(w, r, http.StatusNotFound)
	default:
		errorHandler(w, r, http.StatusNotFound)
	}
}

func messagePostHandler(w http.ResponseWriter, r *http.Request) {
	for k, v := range r.Form {
		fmt.Println("key:", k)
		fmt.Println("val:", strings.Join(v, ""))
	}

	fmt.Printf("%+v\n", r)

	fmt.Fprintf(w, "Thanks for the data!") // write data to response
}

// https://stackoverflow.com/questions/9996767/showing-custom-404-error-page-with-standard-http-package
func errorHandler(w http.ResponseWriter, r *http.Request, status int) {
	w.WriteHeader(status)
	if status == http.StatusNotFound {
		fmt.Fprint(w, "custom 404")
	}
}
