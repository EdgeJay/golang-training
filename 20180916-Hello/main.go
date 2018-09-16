package main

import "fmt"

func main() {
	x := 42

	fmt.Println("Hello World!")
	fmt.Printf("%d in binary is %b\n", x, x)
	fmt.Printf("%d in hex is %#X\n", x, x)

	fmt.Println("Looping through some numbers...")

	for i := 1; i <= 50; i++ {
		fmt.Printf("%d\t%b\t%x\t%q\n", i, i, i, i)
	}
}
