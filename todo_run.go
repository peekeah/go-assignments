package main

import "fmt"

func RunTodos() {
	initialTasks := []RawTask{
		{Description: "Buy groceries (milk, bread, eggs)", IsComplete: false},
		{Description: "Finish reading 'Atomic Habits'", IsComplete: true},
		{Description: "Complete Go language assignment", IsComplete: false},
		{Description: "Call the bank for account update", IsComplete: false},
		{Description: "Pay electricity bill", IsComplete: true},
		{Description: "Book doctor’s appointment", IsComplete: false},
		{Description: "Clean the kitchen", IsComplete: false},
		{Description: "Reply to pending emails", IsComplete: true},
		{Description: "Prepare for team meeting tomorrow", IsComplete: false},
		{Description: "Workout for 30 minutes", IsComplete: true},
	}

	todos := InitilizeTodo()
	todos.AddBulkTasks(initialTasks)
	todos.AddTask(RawTask{"Buy Phone", false})

	taskId := todos.ListIncompleteTasks()[0].Id
	task, err := todos.GetTask(taskId)
	if err != nil {
		fmt.Printf("Task not found for id: %d\n", taskId)
	}

	fmt.Println(todos.ListTasks())
	task.MarkComplete()
	todos.DeleteTask(taskId)

	fmt.Println(todos.ListTasks())
}
