package main

import (
	"errors"
	"math/rand"
)

type Task struct {
	Id          int
	Description string
	IsComplete  bool
}

type RawTask struct {
	Description string
	IsComplete  bool
}

type Todo []*Task

func InitilizeTodo() *Todo {
	return &Todo{}
}

func (t *Todo) AddBulkTasks(tasks []RawTask) {
	if len(tasks) != 0 {
		for _, val := range tasks {
			t.AddTask(val)
		}
	}
}

func (t *Todo) AddTask(task RawTask) Task {
	newId := rand.Intn(90000) + 10000
	newTask := Task{newId, task.Description, task.IsComplete}
	*t = append(*t, &newTask)
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

func (t *Task) MarkComplete() {
	t.IsComplete = true
}

func (t *Todo) DeleteTask(taskId int) error {
	for id, val := range *t {
		if val.Id == taskId {
			*t = append((*t)[:id], (*t)[id+1:]...)
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
