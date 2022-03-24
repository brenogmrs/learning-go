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

	jsonD1, error := json.Marshal(d)

	if error != nil {
		log.Fatal(error)
	}

	fmt.Println(bytes.NewBuffer(jsonD1))

	d2 := map[string]string{
		"nome":  "Lisa",
		"breed": "poodle",
	}

	jsonD2, error := json.Marshal(d2)
	if error != nil {
		log.Fatal(error)
	}

	fmt.Println(bytes.NewBuffer(jsonD2))

}
