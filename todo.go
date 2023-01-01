package todo

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"time"

	"github.com/alexeyco/simpletable"
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

	ls[index-1].Completed = time.Now()
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
	table := simpletable.New()

	table.Header = &simpletable.Header{
		Cells: []*simpletable.Cell{
			{Align: simpletable.AlignCenter, Text: "#"},
			{Align: simpletable.AlignCenter, Text: "Task"},
			{Align: simpletable.AlignCenter, Text: "Done?"},
			{Align: simpletable.AlignRight, Text: "Created"},
			{Align: simpletable.AlignRight, Text: "Completed"},
		},
	}

	var cells [][]*simpletable.Cell

	for idx, ell := range *t {
		idx++

		task := blue(ell.Title)
		done := blue("no")
		completedAt := blue("")
		if ell.Done {
			task = green(fmt.Sprintf("\u2705 %s", ell.Title))
			done = green("yes")
			completedAt = green(ell.Completed.Format(time.RFC822))
		}

		cells = append(cells, []*simpletable.Cell{
			{Text: fmt.Sprintf("%d", idx)},
			{Text: task},
			{Text: done},
			{Text: ell.CreatedAt.Format(time.RFC822)},
			{Text: completedAt},
		})
	}

	ls := *t
	if len(ls) > 0 {
		table.Body = &simpletable.Body{Cells: cells}

		table.Footer = &simpletable.Footer{Cells: []*simpletable.Cell{
			{Align: simpletable.AlignCenter, Span: 5, Text: red(fmt.Sprintf("You need to complete %d tasks", t.CountPending()))},
		}}

		table.SetStyle(simpletable.StyleUnicode)

		table.Println()
	} else {
		fmt.Println(`
        At real time there are not tasks.

        Create new one by command:
            -add <your_task_title> 
        `)
	}

}

func (t *Todos) CountPending() int {
	total := 0
	for _, ell := range *t {
		if !ell.Done {
			total++
		}
	}

	return total
}
