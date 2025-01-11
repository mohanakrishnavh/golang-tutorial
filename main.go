package main

import "fmt"

func main() {
	// Declaring Variables
	var appName = "MyTodoListApp"
	// Slice - Dynamic, Array - Fixed Size
	var tasks []string
	// var limit = 20

	fmt.Println("###### Welcome to Todolist App! :", appName, "######")
	fmt.Println("### Tasks ###")

	tasks = addTask(tasks, "Watch Go crash course")
	tasks = addTask(tasks, "Watch Go full course")
	tasks = addTask(tasks, "Reward myself with a nice meal")
	printTasks(tasks)
}

func addTask(tasks []string, task string) []string {
	tasks = append(tasks, task)
	return tasks
}

func printTasks(tasks []string) {
	// Loops
	for idx, task := range tasks {
		fmt.Printf("%d. %s\n", idx+1, task)
	}
}
