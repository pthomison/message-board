package main

import (
	"embed"
	"html/template"
	"net/http"

	utils "github.com/pthomison/golang-utils"
)

//go:embed ui/*.css
var css embed.FS

func cssHandler(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFS(css, "ui/*.css")
	utils.Check(err)
	err = t.Execute(w, nil)
	utils.Check(err)
}
