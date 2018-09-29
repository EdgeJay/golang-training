package main

import "fmt"

type person struct {
	name   string
	age    int
	gender string
}

func main() {
	joe := person{"Joe", 25, "male"}
	fmt.Printf("%T, %v\n", joe, joe)

	pointer := &person{"Jane", 24, "female"}
	fmt.Printf("%T, %v\n", pointer, pointer)

	// modify actual struct through pointer
	(*pointer).name = "JaneJane"
	(*pointer).age = 26
	fmt.Printf("%T, %v\n", pointer, pointer)

	jenny := person{"Jenny", 14, "female"}
	pointer = &jenny
	fmt.Printf("%T, %v\n", pointer, pointer)
	(*pointer).name = "JenJen"
	fmt.Printf("%T, %v\n", pointer, pointer)
}
