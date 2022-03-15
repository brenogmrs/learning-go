package main

import "fmt"

func main() {

	fmt.Println("Pointers")

	var v1 int = 10
	var v2 int = v1

	fmt.Println(v1, v2)

	v1++

	fmt.Println(v1, v2)

	// A pointer is a memory reference

	var v3 int = 100
	var pointer *int

	pointer = &v3

	fmt.Println(v3, *pointer)

	v3++

	fmt.Println(v3, *pointer)
}
