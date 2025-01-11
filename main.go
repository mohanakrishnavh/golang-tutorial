package main

import (
	"fmt"
	"net/http"
)

var task1 = "Finish the Go crash course"
var task2 = "Finish the Go Full course"
var task3 = "Go for a run"
var tasks = []string{task1, task2, task3}

func main() {
	fmt.Println("###### Welcome to Todolist App! ######")

	http.HandleFunc("/", handleUser)
	http.HandleFunc("/show-tasks", showTasks)

	http.ListenAndServe("localhost:8080", nil)
}

func handleUser(writer http.ResponseWriter, request *http.Request) {
	var greeting = "Hello user. Weclome to TodoList App!"
	fmt.Fprintf(writer, greeting)
}

func showTasks(writer http.ResponseWriter, request *http.Request) {
	for _, task := range tasks {
		fmt.Fprintln(writer, task)
	}
}

func addTask(tasks []string, task string) []string {
	tasks = append(tasks, task)
	return tasks
}

func printTasks(tasks []string) {
	fmt.Println("### Tasks ###")
	// Loops
	for idx, task := range tasks {
		fmt.Printf("%d. %s\n", idx+1, task)
	}
}
