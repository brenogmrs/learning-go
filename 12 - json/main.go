package main

import (
	"bytes"
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
	d := dog{"Rex", "Dalmata", 2}
	fmt.Println(d)

	json, error := json.Marshal(d)

	if error != nil {
		log.Fatal(error)
	}

	fmt.Println(bytes.NewBuffer(json))

}
