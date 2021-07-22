// main.go
package main

import (
  "embed"

  "github.com/pthomison/message-board/cmd"
)

//go:embed ui/*
var uiAssests embed.FS

func main() {
  println("Ba dum, tss!")

  cmd.RunServer(uiAssests)
}
