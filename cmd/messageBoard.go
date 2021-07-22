package cmd

import (
	"html/template"
	"net/http"

	utils "github.com/pthomison/golang-utils"
)

type MessageBoard struct {
	Messages []Message
}

// //go:embed ui/message-board.html
// var mbHtml embed.FS

func (s *server) messageBoardHandler(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFS(s.uiAssests, "ui/message-board.html")
	utils.Check(err)
	err = t.Execute(w, s.mb)
	utils.Check(err)
}
