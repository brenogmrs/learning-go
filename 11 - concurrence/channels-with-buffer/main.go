package main

import "fmt"

func main() {

	ch := make(chan string, 3)

	ch <- "Hello World"

	message := <-ch

	fmt.Println(message)
}
