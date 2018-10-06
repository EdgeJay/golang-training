package main

import "fmt"

type person struct {
	name   string
	age    int
	gender string
}

func (p person) tryChangeName(name string) {
	p.name = name
}

func (p *person) reallyChangeName(name string) {
	p.name = name
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

	mark := &person{"Mark", 30, "male"}
	mark.tryChangeName("Marky")
	fmt.Printf("Mark try to change name to: %s, got: %s\n", "Marky", mark.name)
	mark.reallyChangeName("Marko")
	fmt.Printf("Mark try to change name to: %s, got: %s\n", "Marko", mark.name)

	sean := person{"Sean", 30, "male"}
	sean.tryChangeName("Seano")
	fmt.Printf("Sean try to change name to: %s, got: %s\n", "Seano", sean.name)
	sean.reallyChangeName("Seaaan")
	fmt.Printf("Sean try to change name to: %s, got: %s\n", "Seaaan", sean.name)
}
