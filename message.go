package main

type Message struct {
	Author  string
	Message string
	Date    string
}

// func messageBoardHandler(w http.ResponseWriter, r *http.Request) {
// 	t, err := template.ParseFS(mbHtml, "ui/*")
// 	utils.Check(err)
// 	err = t.Execute(w, nil)
// 	utils.Check(err)
// }
