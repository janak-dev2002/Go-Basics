package main

import (
	"fmt"
	"net/http"
)

var shortGolang = "Watch Go crash course"
var fullGolang = "Watch Nana's Golang Full Course"
var rewardDessert = "Reward myself with a donut"
var taskItems = []string{shortGolang, fullGolang, rewardDessert}

// Handler function for the root path.
func helloHandler(writer http.ResponseWriter, request *http.Request) {
	a, err := fmt.Fprintln(writer, "Hello,World!") // This returns number of bytes written and an error
	fmt.Println(a, err)
}

func taskHandler(writer http.ResponseWriter, request *http.Request) {
	for index, task := range taskItems{
		a,err := fmt.Fprintf(writer, "Task %d: %s\n", index+1, task)
		fmt.Println(a, err)
	}
}

func main() {
	fmt.Println("Welcome to the Web API server on port 8080")
	http.HandleFunc("/hello", helloHandler)
	http.HandleFunc("/show-task",taskHandler)
	http.ListenAndServe(":8080", nil)
}

