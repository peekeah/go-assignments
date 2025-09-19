package main

import "fmt"

func RunTodos() {
	initialTasks := []RawTask{
		{Description: "Buy groceries (milk, bread, eggs)", IsComplete: false},
		{Description: "Finish reading 'Atomic Habits'", IsComplete: true},
		{Description: "Complete Go language assignment", IsComplete: false},
		{Description: "Call the bank for account update", IsComplete: false},
		{Description: "Pay electricity bill", IsComplete: true},
	}

	todos := InitializeTodo()
	todos.AddBulkTasks(initialTasks)           // Bulk upload
	todos.AddTask(RawTask{"Buy Phone", false}) // Add task

	taskId := todos.ListIncompleteTasks()[0].Id
	task, err := todos.GetTask(taskId)
	if err != nil {
		fmt.Printf("Task not found for id: %d\n", taskId)
	}

	fmt.Println(todos.ListTasks())
	todos.MarkComplete(task) // Complete
	todos.DeleteTask(taskId) // Delete

	fmt.Println(todos.ListTasks()) // List all task

	// Test file Sync
	newTask := todos.AddTask(RawTask{"Hair Cut", false})
	todos.MarkComplete(newTask)
	tt, err := todos.GetTask(newTask.Id)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(tt)
}
