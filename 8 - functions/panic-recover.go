package main

import "fmt"

func recoverExecution() {
	fmt.Println("trying to recover execution")
	if recovery := recover(); recovery != nil {
		fmt.Println("recorery executed successfully")
	}
}

func isStudentAprooved(n1, n2 float64) bool {
	defer recoverExecution()
	avg := (n1 + n2) / 2

	if avg > 6 {
		return true
	} else if avg < 6 {
		return false
	}

	panic("The Average cannot be 6")
}

func main() {
	fmt.Println("Panic! and Recover")

	fmt.Println(isStudentAprooved(6, 6))
	fmt.Println("executed!")
}
