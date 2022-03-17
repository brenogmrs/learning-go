package main

import "fmt"

type user struct {
	name string
	age  uint8
}

func (u user) save() {
	fmt.Printf("Saving user %s data in the database \n", u.name)
}

func (u user) isAdult() bool {
	return u.age >= 18
}

func (u *user) birthday() {
	u.age++
}

func main() {
	fmt.Println("Methods")

	user := user{"Breno", 23}

	user.save()

	isUserAdult := user.isAdult()

	fmt.Println(isUserAdult)

	user.birthday()
	fmt.Println(user.age)
}
