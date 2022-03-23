package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	ch := multiplex(write("Hello World"), write("Coding in Go!"))

	var messages []string

	for i := 0; i < 10; i++ {
		fmt.Println(<-ch)
		messages = append(messages, <-ch)
	}

	fmt.Println(len(messages))

}

func multiplex(channelInput1, channelInput2 <-chan string) <-chan string {

	outputChannel := make(chan string)

	go func() {
		for {
			select {
			case message := <-channelInput1:
				outputChannel <- message
			case message := <-channelInput2:
				outputChannel <- message
			}
		}
	}()

	return outputChannel
}

func write(text string) <-chan string {

	ch := make(chan string)

	go func() {
		for {
			ch <- fmt.Sprintf("Received value %s", text)
			time.Sleep(time.Millisecond * time.Duration(rand.Intn(2000)))
		}
	}()

	return ch
}
