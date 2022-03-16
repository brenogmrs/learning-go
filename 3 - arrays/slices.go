package main

import (
	"fmt"
)

func main() {

	slice := []int{10, 20, 30, 40}
	anotherSlice := []int{50, 60, 70, 80}

	slice = append(slice, anotherSlice...)

	//Teste(slice...)

	//SlicingASlice()

	// internal arrays

	slice2 := make([]float32, 10, 11)
	fmt.Println(slice2)
	fmt.Println(len(slice2)) // length
	fmt.Println(cap(slice2)) // capacity

	slice2 = append(slice2, 2)
	fmt.Println(slice2)
	slice2 = append(slice2, 2)
	fmt.Println(cap(slice2))
}
func Teste(inteiros ...int) {
	total := 0

	for i, val := range inteiros {
		fmt.Println("At index", i, "we have the value", val)
		total += val
	}
}

func SlicingASlice() {
	//                     0             1           2          3              4
	flavors := []string{"pepperoni", "mozzarella", "marg", "pineapple", "calabrezza"}

	//deletando valor de uma slice
	flavors = append(flavors[:3], flavors[3:]...)

	// acessando até o item 2 (1 off)
	slice := flavors[:3]

	// acessando do item 3 até o final
	slice = flavors[3:]

	fmt.Println(slice)
	fmt.Println(flavors)
}
