package main

import "fmt"

func f1() {
	fmt.Println("executing function 1")
}

func f2() {
	fmt.Println("executing function 2")
}

func isStudentAprooved(n1, n2 float32) bool {
	defer fmt.Println("Average calculated. Returning result")
	fmt.Println("Verifying if the student is aprooved")

	m := (n1 + n2) / 2

	if m >= 6 {
		return true
	}

	return false
}

func main() {
	// defer = adiar
	fmt.Println("Defer clause")
	fmt.Println(isStudentAprooved(1, 8))
}
