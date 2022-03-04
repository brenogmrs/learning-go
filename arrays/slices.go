package main

import (
	"fmt"
)

func main() {

	slice := []int{10, 20, 30, 40}

	slice = append(slice, 50)

	total := 0

	for i, val := range slice {
		fmt.Println("At index", i, "we have the value", val)
		total += val
	}

	fmt.Println(total)

}
