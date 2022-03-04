package main

import (
	"fmt"
)

func main() {

	slice := []int{10, 20, 30, 40}
	anotherSlice := []int{50, 60, 70, 80}

	slice = append(slice, anotherSlice...)

	total := 0

	for i, val := range slice {
		fmt.Println("At index", i, "we have the value", val)
		total += val
	}

	fmt.Println(total)

}
