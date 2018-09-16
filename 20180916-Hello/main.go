package main

import "fmt"

func main() {
	x := 42
	fmt.Println("Hello World!")
	fmt.Printf("%d in binary is %b\n", x, x)
	fmt.Printf("%d in hex is %#X\n", x, x)
}
