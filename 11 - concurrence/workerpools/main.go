package main

import (
	"fmt"
)

func main() {
	fmt.Println("Recusive Functions speed up by channels")

	tasks := make(chan int, 45)
	results := make(chan int, 45)

	go worker(tasks, results)
	go worker(tasks, results)

	for i := 0; i < 45; i++ {
		tasks <- i
	}

	close(tasks)

	for i := 0; i < 45; i++ {
		result := <-results
		fmt.Println(result)
	}

}

func worker(tasks <-chan int, results chan<- int) {
	for task := range tasks {
		results <- fibbo(task)
	}
}

func fibbo(position int) int {
	if position <= 1 {
		return position
	}

	return fibbo(position-2) + fibbo(position-1)
}
