package main

import (
	"fmt"
	"time"
)

func main() {

	ch := write("hello world")

	for i := 0; i < 10; i++ {
		fmt.Println(<-ch)
	}
}

func write(text string) <-chan string {

	ch := make(chan string)

	go func() {
		for {
			ch <- fmt.Sprintf("Received value %s", text)
			time.Sleep(time.Millisecond * 500)
		}
	}()

	return ch
}
