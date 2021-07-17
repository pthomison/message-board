package main

import (
	"embed"
	"html/template"
	"net/http"

	utils "github.com/pthomison/golang-utils"
)

type MessageBoard struct {
	Messages []Message
}

//go:embed ui/message-board.html
var mbHtml embed.FS

func messageBoardHandler(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFS(mbHtml, "ui/*")
	utils.Check(err)
	err = t.Execute(w, bones())
	utils.Check(err)
}
