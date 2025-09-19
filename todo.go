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

func InitializeTodo() *Todo {
	data, err := readJson()
	if err != nil {
		fmt.Println(err)
	}

	return &data
}

func saveJson(fileContent Todo) error {
	// Save to JSON
	data, err := json.Marshal(fileContent)
	if err != nil {
		return errors.New("error while storing data")
	}

	return os.WriteFile("todo.json", data, 0644)
}

func readJson() (Todo, error) {
	data, err := os.ReadFile("todo.json")
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return Todo{}, nil
		}
		return nil, fmt.Errorf("error while reading data: %w", err)
	}

	todos := Todo{}
	if err := json.Unmarshal(data, &todos); err != nil {
		return nil, fmt.Errorf("error while unmarshaling json: %w", err)
	}

	return todos, nil
}

func (t *Todo) SyncFile() {
	if err := saveJson(*t); err != nil {
		fmt.Println("failed to sync file:", err)
	}
}

func (t *Todo) AddBulkTasks(tasks []RawTask) {
	if len(tasks) != 0 {
		for _, val := range tasks {
			t.AddTask(val, false)
		}
	}
	t.SyncFile()
}

func (t *Todo) AddTask(task RawTask, enableSync ...bool) *Task {
	newId := rand.Intn(90000) + 10000

	newTask := &Task{newId, task.Description, task.IsComplete}
	*t = append(*t, newTask)

	doSync := true
	if len(enableSync) > 0 {
		doSync = enableSync[0]
	}

	if doSync {
		t.SyncFile()
	}

	return newTask
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
			t.SyncFile()
			return nil
		}
	}
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
