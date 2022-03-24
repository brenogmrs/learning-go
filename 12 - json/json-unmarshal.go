package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type dog struct {
	Name  string `json: "name"`
	Breed string `json: "breed"`
	Age   uint   `json: "age"`
}

func main() {
	dJSON := `{"breed":"poodle","name":"Lisa", "age": 10}`

	var d dog

	if err := json.Unmarshal([]byte(dJSON), &d); err != nil {
		log.Fatal(err)
	}

	fmt.Println(d)

	d2JSON := `{"Name":"Rex","Breed":"Dalmata"}`

	var d2 = make(map[string]string)

	if err := json.Unmarshal([]byte(d2JSON), &d2); err != nil {
		log.Fatal(err)
	}

	fmt.Println(d2)

}
