package cmd

import (
	"fmt"
	"net/http"
	"time"
)

type Message struct {
	Author    string
	Message   string
	Date      string
	Timestamp time.Time
}

func (s *server) messageHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		s.errorHandler(w, r, http.StatusNotFound)
	case http.MethodPost:
		s.messagePostHandler(w, r)
	case http.MethodPut:
		s.errorHandler(w, r, http.StatusNotFound)
	case http.MethodDelete:
		s.errorHandler(w, r, http.StatusNotFound)
	default:
		s.errorHandler(w, r, http.StatusNotFound)
	}
}

// https://stackoverflow.com/questions/9996767/showing-custom-404-error-page-with-standard-http-package
func (s *server) errorHandler(w http.ResponseWriter, r *http.Request, status int) {
	w.WriteHeader(status)
	if status == http.StatusNotFound {
		fmt.Fprint(w, "custom 404")
	}
}

func (s *server) messagePostHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	email := r.Form["email"][0]
	message := r.Form["message"][0]
	timestamp := time.Now()

	fmt.Printf("EMAIL: %+v\n", email)
	fmt.Printf("MESSAGE: %+v\n", message)

	s.addMessage(Message{
		Author:    email,
		Message:   message,
		Timestamp: timestamp,
	})

	fmt.Printf("Post Complete: %+v\n", message)

	http.Redirect(w, r, "/", 301)

	// fmt.Fprintf(w, "Thanks for the data!") // write data to response
}

func (s *server) addMessage(m Message) {
	s.mc <- m
	s.mb.Messages = append(s.mb.Messages, m)
}
