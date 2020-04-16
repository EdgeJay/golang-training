package main

import "fmt"

func sum(input chan int, done chan<- int) {
	acc := 0
	for {
		num, more := <-input
		if more {
			acc += num
		} else {
			done <- acc
			return
		}
	}
}

func main() {
	numbers := []int{1, 2, 3, 5, 8, 13, 21, 34}
	input := make(chan int)
	done := make(chan int)

	go sum(input, done)

	for _, num := range numbers {
		input <- num
	}
	close(input)

	fmt.Printf("Sum: %v\n", <-done)
}
