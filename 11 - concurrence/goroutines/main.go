package main

import (
	"fmt"
	"time"
)

func main() {
	// cocorrĂȘncia != paralelismo
	go write("Hello World!")
	write("coding in Go!")
}

func write(text string) {
	for {
		fmt.Println(text)
		time.Sleep(time.Second)
	}
}
