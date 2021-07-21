package main

import (
	"embed"
	"html/template"
	"net/http"

	utils "github.com/pthomison/golang-utils"
)

//go:embed ui/*.css
var css embed.FS

//go:embed ui/*.js
var js embed.FS

func (s *server) cssHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/css")

	t, err := template.ParseFS(css, "ui/*.css")
	utils.Check(err)
	err = t.Execute(w, nil)
	utils.Check(err)
}

func (s *server) jsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/javascript")

	t, err := template.ParseFS(js, "ui/*.js")
	utils.Check(err)
	err = t.Execute(w, nil)
	utils.Check(err)
}
