package main

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"os"
	"time"
)

type task struct {
	Title     string
	CreatedAt time.Time
	Completed time.Time
	Done      bool
}

type Todos []task

func (t *Todos) Add(task string) {
	todo := task{
		Title:     task,
		CreatedAt: time.Now(),
		Completed: time.Time{},
		Done:      false,
	}
	*t = append(*t, todo)
}

func (t *Todos) Completed(index int) error {
	ls := *t
	if index <= 0 || index > len(ls) {
		return errors.New("invalid index")
	}

	ls[index-1].CreatedAt = time.Now()
	ls[index-1].Done = true

	return nil
}

func (t *Todos) Delete(index int) error {
	ls := *t
	if index <= 0 || index > len(ls) {
		return errors.New("invalid index")
	}

	*t = append(ls[:index-1], ls[index:]...)

	return nil
}

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

func (t *Todos) Store(filename string) error {
	data, err := json.Marshal(t)
	if err != nil {
		return err
	}

	return ioutil.WriteFile(filename, data, 0644)
}
