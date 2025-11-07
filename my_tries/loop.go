package main

import "fmt"

func main() {
	var shortGolang = "Watch Go crash course"
	var fullGolang = "Watch Nana's Golang Full Course"
	var rewardDessert = "Reward myself with a donut"

	var taskItems = []string{shortGolang, fullGolang, rewardDessert}

	fmt.Println("###### Welcome to our Todolist App! ######")
	fmt.Println("List of my Todos")
	fmt.Println()

	// Using for loop to iterate over the slice, when using _ to ignore the index
	for _, task := range taskItems {
		fmt.Println(task)
	}

	fmt.Println()
	fmt.Println("#####*******************######")
	fmt.Println()

	// Using for loop with index to iterate over the slice
	for index, task := range taskItems {
		fmt.Printf("Task %d: %s\n", index+1, task)
	}

	// Using traditional for loop with index
	fmt.Println()
	fmt.Println("#####------------Using traditional for loop with index------------######")
	fmt.Println()
	for i := 0; i < len(taskItems); i++ {
		fmt.Printf("Task %d: %s\n", i+1, taskItems[i])
	}

	// While like for loop

	fmt.Println()
	fmt.Println("#####------------Using While like for loop------------######")
	fmt.Println()

	index := 0

	for index < len(taskItems) {
		fmt.Printf("Task %d: %s\n", index+1, taskItems[index])
		index++
	}

	// Infinite loop (no initialization, no condition, no post)

	fmt.Println()
	fmt.Println("#####------------Infinite loop------------######")
	fmt.Println()
	for {
		fmt.Println("This is an infinite loop. Press Ctrl + C to stop.")
		break
	}

}
