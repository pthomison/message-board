package cmd

import (
	"fmt"
	"html"
	"net/http"
)

func messageBoardHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
}
