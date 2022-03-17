package main

import "fmt"

func closure() func() {
	text := "Text inside the closure function"

	function := func() {
		fmt.Println(text)
	}

	return function
}

func main() {
	fmt.Println("Closure Function")

	newFunction := closure()

	newFunction()
}
