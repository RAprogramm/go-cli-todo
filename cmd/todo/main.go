// Copyright © 2022 RAprogramm <andrey.rozanov.vl@gmail.com>
package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/RAprogramm/go-cli-todo"
)

const (
	todoFile = ".todos.json"
)

// main function runs app
func main() {
	todos := &todo.Todos{}

	if err := todos.Load(todoFile); err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}
	todos.Print()
Loop:
	for {
		switch todo.Select() {
		case "Create a new task":
			newTask := todo.Input("What do you need to do?")
			todos.Add(newTask)
			err := todos.Store(todoFile)
			if err != nil {
				fmt.Fprintln(os.Stderr, err.Error())
				os.Exit(1)
			}
			todos.Print()
		case "Mark task as completed":
			completedTask, _ := strconv.Atoi(todo.Input("What task did you complete?"))
			err := todos.Completed(completedTask)
			if err != nil {
				fmt.Fprintln(os.Stderr, err.Error())
				os.Exit(1)
			}
			err = todos.Store(todoFile)
			if err != nil {
				fmt.Fprintln(os.Stderr, err.Error())
				os.Exit(1)
			}
			todos.Print()
		case "Delete a task":
			delTask, _ := strconv.Atoi(todo.Input("What task you want to delete?"))
			err := todos.Delete(delTask)
			if err != nil {
				fmt.Fprintln(os.Stderr, err.Error())
				os.Exit(1)
			}
			err = todos.Store(todoFile)
			if err != nil {
				fmt.Fprintln(os.Stderr, err.Error())
				os.Exit(1)
			}
			todos.Print()
		case "Exit":
			os.Exit(0)
			break Loop
		}
	}
}
