package todo

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"time"
)

type Task struct {
	Title     string
	CreatedAt time.Time
	Completed time.Time
	Done      bool
}

type Todos []Task

// Add method create a task and add it to task list
func (t *Todos) Add(task string) {
	todo := Task{
		Title:     task,
		CreatedAt: time.Now(),
		Completed: time.Time{},
		Done:      false,
	}
	*t = append(*t, todo)
}

// Completed method change index of a task to mark it as completed
func (t *Todos) Completed(index int) error {
	ls := *t
	if index <= 0 || index > len(ls) {
		return errors.New("invalid index")
	}

	ls[index-1].CreatedAt = time.Now()
	ls[index-1].Done = true

	return nil
}

// Delete method remove task from list
func (t *Todos) Delete(index int) error {
	ls := *t
	if index <= 0 || index > len(ls) {
		return errors.New("invalid index")
	}

	*t = append(ls[:index-1], ls[index:]...)

	return nil
}

// Load method read json file and return it
func (t *Todos) Load(filename string) error {
	file, err := ioutil.ReadFile(filename)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return nil
		}
		return err
	}

	if len(file) == 0 {
		return err
	}
	err = json.Unmarshal(file, t)
	if err != nil {
		return err
	}

	return nil
}

// Store method save task into json file
func (t *Todos) Store(filename string) error {
	data, err := json.Marshal(t)
	if err != nil {
		return err
	}

	return ioutil.WriteFile(filename, data, 0644)
}

// Print method get all tasks from file
func (t *Todos) Print() {
	for i, item := range *t {
		i++
		fmt.Printf("%d - %s\n", i, item.Title)
	}
}
