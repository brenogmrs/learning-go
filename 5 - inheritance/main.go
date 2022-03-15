package main

import (
	"fmt"
)

type person struct {
	name    string
	surname string
	age     uint8
	height  uint8
}

type student struct {
	person
	course     string
	university string
}

func main() {

	fmt.Println("{Inheritance}")

	p1 := person{"Breno", "Guimarães", 23, 181}

	s1 := student{p1, "I.T", "MIT"}

	fmt.Println(s1.name)
}
