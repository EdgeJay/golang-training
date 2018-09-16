package main

import "fmt"

func main() {
	for i := 1; i <= 25; i++ {
		fmt.Printf("Count is %d\n", i)
	}

	// for is Go's while
	fmt.Println("Test another kind of for-loop...")
	j := 0
	for j < 50 {
		j++
		fmt.Printf("Count is %d\n", j)
	}
}
