package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"math/rand"
	"os"
)

type Task struct {
	Id          int    `json:"id"`
	Description string `json:"description"`
	IsComplete  bool   `json:"is_complete"`
}

type RawTask struct {
	Description string
	IsComplete  bool
}

type Todo []*Task

func InitilizeTodo() *Todo {
	data, err := readJson()
	if err != nil {
		fmt.Println(err)
	}

	ptrArr := returnPtrArr(data)
	return &ptrArr
}

func returnRawArr(todos Todo) []Task {
	res := []Task{}

	for _, val := range todos {
		res = append(res, *val)
	}
	return res
}

func returnPtrArr(todos []Task) Todo {
	res := Todo{}
	for id, val := range todos {
		res = append(res, &val)
		res[id] = &val
	}
	return res
}

func saveJson(fileContent []Task) error {
	// Save to JSON
	data, err := json.Marshal(fileContent)
	if err != nil {
		return errors.New("error while storing data")
	}

	os.WriteFile("todo.json", data, 0644)
	return nil
}

func readJson() ([]Task, error) {
	data, err := os.ReadFile("todo.json")
	if err != nil {
		return nil, errors.New("error while storing data")
	}

	todos := []Task{}
	json.Unmarshal(data, &todos)

	return todos, nil
}

func (t *Todo) SyncFile() {
	content := returnRawArr(*t)
	saveJson(content)
}

func (t *Todo) AddBulkTasks(tasks []RawTask) {
	if len(tasks) != 0 {
		for _, val := range tasks {
			t.AddTask(val)
		}
	}
	t.SyncFile()
}

func (t *Todo) AddTask(task RawTask) *Task {
	newId := rand.Intn(90000) + 10000
	newTask := Task{newId, task.Description, task.IsComplete}
	*t = append(*t, &newTask)
	t.SyncFile()
	return (*t)[len(*t)-1]
}

func (t *Todo) GetTask(taskId int) (*Task, error) {
	for _, val := range *t {
		if val.Id == taskId {
			return val, nil
		}
	}
	return nil, errors.New("task not found")
}

func (t *Todo) MarkComplete(task *Task) {
	task.IsComplete = true
	t.SyncFile()
}

func (t *Todo) DeleteTask(taskId int) error {
	for id, val := range *t {
		if val.Id == taskId {
			*t = append((*t)[:id], (*t)[id+1:]...)
			return nil
		}
	}
	t.SyncFile()
	return errors.New("task not found")
}

func (t *Todo) ListTasks() (res []Task) {
	for _, val := range *t {
		res = append(res, *val)
	}
	return res
}

func (t *Todo) ListIncompleteTasks() (res []Task) {
	for _, val := range *t {
		if !val.IsComplete {
			res = append(res, *val)
		}
	}

	return
}
