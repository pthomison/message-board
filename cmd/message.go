package cmd

import (
	"fmt"
	"net/http"
)

type Message struct {
	Author  string
	Message string
	Date    string
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

func (s *server) messagePostHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	// for k, v := range r.Form {
	// 	fmt.Println("key:", k)
	// 	fmt.Println("val:", strings.Join(v, ""))
	// }

	email := r.Form["email"][0]
	message := r.Form["message"][0]

	fmt.Printf("EMAIL: %+v\n", email)
	fmt.Printf("MESSAGE: %+v\n", message)

	s.mb.Messages = append(s.mb.Messages, Message{
		Author:  email,
		Message: message,
	})

	http.Redirect(w, r, "/", 301)

	// fmt.Fprintf(w, "Thanks for the data!") // write data to response
}

// https://stackoverflow.com/questions/9996767/showing-custom-404-error-page-with-standard-http-package
func (s *server) errorHandler(w http.ResponseWriter, r *http.Request, status int) {
	w.WriteHeader(status)
	if status == http.StatusNotFound {
		fmt.Fprint(w, "custom 404")
	}
}
