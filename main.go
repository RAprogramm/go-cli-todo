/*
Copyright © 2022 RAprogramm <andrey.rozanov.vl@gmail.com>
*/
package main

import (
	"flag"

	"github.com/RAprogramm/go-cli-todo/cmd"
)

const todoFile = ".todos.json"

// main function runs app
func main() {
	cmd.Execute()
	add := flag.Bool("add", false, "add new task")
	flag.Parse()
	todos := &task.Todos{}
}
