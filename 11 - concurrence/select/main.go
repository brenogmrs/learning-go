package main

import (
	"fmt"
	"time"
)

func main() {
	ch1, ch2 := make(chan string), make(chan string)

	go func() {
		for {
			time.Sleep(time.Millisecond * 500)
			ch1 <- "channel one"
		}
	}()

	go func() {
		for {
			time.Sleep(time.Second * 2)
			ch2 <- "channel two"
		}
	}()

	for {
		select {
		case messageCh1 := <-ch1:
			fmt.Println(messageCh1)
		case messageCh2 := <-ch2:
			fmt.Println(messageCh2)
		}
	}

}
