package cmd

import (
	"html/template"
	"net/http"

	utils "github.com/pthomison/golang-utils"
	"golang.org/x/net/websocket"
)

type MessageBoard struct {
	Messages    []Message
	Connections map[*websocket.Conn]chan Message
}

func (s *server) messageBoardHandler(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFS(s.uiAssests, "ui/message-board.html")
	utils.Check(err)
	err = t.Execute(w, s.mb)
	utils.Check(err)
}
