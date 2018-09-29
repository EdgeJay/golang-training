package main

import "fmt"

type foo struct {
	bar string
}

/**
 * Test how variables are initialized with default values when no values are provided
 */
func main() {
	var aInt int
	var aString string
	var aSlice []int
	var aArray [5]string
	var aFoo foo

	fmt.Printf("%T, value: %d\n", aInt, aInt)
	fmt.Printf("%T, value: %s, length: %d\n", aString, aString, len(aString))
	fmt.Printf("%T, length: %d\n", aSlice, len(aSlice))
	fmt.Printf("%T, length: %d\n", aArray, len(aArray))
	fmt.Printf("%T, value: %s, length: %d\n", aArray[0], aArray[0], len(aArray[0]))
	fmt.Printf("%T, value: %v\n", aFoo, aFoo)
	fmt.Printf("%T, value: %v\n", aFoo.bar, aFoo.bar)
}
