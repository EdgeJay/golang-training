package main

import "fmt"

func sum(ch chan<- int, nums ...int) {
	s := 0
	for _, n := range nums {
		s += n
	}
	ch <- s
}

func main() {
	ch := make(chan int)
	numbers := []int{1, 2, 3, 5, 8, 13, 21, 34}
	sums := []int{}

	for index := 1; index <= len(numbers); index++ {
		arr := numbers[:index]
		go sum(ch, arr...)
		sums = append(sums, <-ch)
	}

	fmt.Printf("%v\n", sums)
	fmt.Println("done")
}
