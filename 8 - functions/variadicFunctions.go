package main

import "fmt"

func sumN(numbers ...int) int {

	total := 0

	for _, num := range numbers {
		total += num
	}

	return total

}

func write(text string, numbers ...int) {

	for _, num := range numbers {
		fmt.Println(text, num)
	}
}

func main() {
	fmt.Println("Variadic Functions")

	total := sumN(1, 2, 3, 4, 5)

	fmt.Println(total)

	write("Hello World", 10, 10, 10)

}
