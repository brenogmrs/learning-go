package main

import "fmt"

func calcSumSub(n1, n2 int) (sum int, sub int) {

	sum = n1 + n2
	sub = n1 - n2

	return
}

func main() {
	fmt.Println("Functions with named return")

	sum, sub := calcSumSub(10, 20)

	fmt.Println(sum, sub)
}
