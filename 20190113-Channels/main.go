package main

import "fmt"

func main() {
	// channels by default are "unbuffred", the following
	// creates a "buffered" channel of strings that is
	// buffered up to 2 values.
	// Buffered channels accept a limited number of values without
	// a corresponding receiver for those values
	// Extracted from: https://gobyexample.com/channel-buffering
	ch := make(chan string, 2)

	ch <- "Hello"
	ch <- "World"

	fmt.Println(<-ch)
	fmt.Println(<-ch)
}
