package main

import (
	"fmt"
	"time"
)

func main() {

	ch := make(chan string)

	go write("Hellow World", ch)

	fmt.Println("Depois da função write começar a ser executada")

	for message := range ch {
		fmt.Println(message)
	}

	fmt.Println("End of Program")

}

func write(text string, channel chan string) {
	for i := 0; i < 5; i++ {
		channel <- text

		time.Sleep(time.Second)
	}

	close(channel)
}
