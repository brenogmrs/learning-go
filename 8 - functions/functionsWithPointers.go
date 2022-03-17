package main

import (
	"fmt"
)

func invertNumber(number int) int {
	return number * -1
}

func inverseNumberWithPointer(number *int) int {
	*number = *number * -1

	return *number
}

func main() {
	fmt.Println("Functions with Pointers")

	number := 20
	inverseNumber := inverseNumberWithPointer(&number)

	fmt.Println(inverseNumber)
	fmt.Println(number)
}
