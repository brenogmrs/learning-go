package main

import (
	"fmt"
)

type user struct {
	name    string
	age     uint8
	address address
}

type address struct {
	street string
	number uint8
}

func main() {

	fmt.Println("{Structs}")

	var u user
	u.name = "Breno"
	u.age = 23

	addr := address{"fools street", 0}

	fmt.Println(u)

	u2 := user{"Breno", 23, addr}
	fmt.Println(u2)

	u3 := user{name: "breno"}
	fmt.Println(u3)
}
