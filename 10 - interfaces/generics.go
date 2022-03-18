package main

import "fmt"

func generic(i interface{}) {
	fmt.Println(i)
}

func main() {
	fmt.Println("Generic types")
	generic("string")
	generic(1)
	generic(true)

	m := map[interface{}]interface{}{
		1:            "string",
		float32(200): true,
		"string":     "string",
		true:         float64(10),
	}

	fmt.Println(m)
}
