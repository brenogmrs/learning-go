package main

import "fmt"

func main() {
	fmt.Println("Maps")

	user := map[string]string{
		"name":    "Breno",
		"surname": "Guimarães",
	}

	fmt.Println(user["name"])

	user2 := map[string]map[string]string{
		"info": {
			"name": "Breno",
		},
		"course": {
			"name":   "engeneering",
			"campus": "campus 1",
		},
	}

	fmt.Println(user2)

	// removing a key from map

	delete(user2, "info")

	fmt.Println(user2)

}
