package main

import (
	"fmt"
	"time"
)

func main() {
	// cocorrência != paralelismo
	go write("Hello World!")
	write("coding in Go!")
}

func write(text string) {
	for {
		fmt.Println(text)
		time.Sleep(time.Second)
	}
}
