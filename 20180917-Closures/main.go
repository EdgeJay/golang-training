package main

import "fmt"

func increment(x int) func() int {
	y := x
	return func() int {
		y++
		return y
	}
}

func main() {
	x := 0
	inc := func() int {
		x++
		return x
	}

	fmt.Println(inc())
	fmt.Println(inc())
	fmt.Println(inc())

	inc2 := increment(x)
	fmt.Println(inc2())
	fmt.Println(inc2())
}
