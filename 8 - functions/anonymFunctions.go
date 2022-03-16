package main

import "fmt"

func main() {
	fmt.Println("Anonym Functions")

	anonymFunctionReturn := func(text string) string {
		return fmt.Sprintf("Received -> %s", text)
	}("Param")

	fmt.Println(anonymFunctionReturn)
}
