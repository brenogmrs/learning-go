package main

import "fmt"

func fibbo(position uint) uint {
	if position <= 1 {
		return position
	}

	return fibbo(position-2) + fibbo(position-1)
}

func main() {
	fmt.Println("Recusive Functions")

	result := fibbo(10)

	fmt.Println(result)
}
