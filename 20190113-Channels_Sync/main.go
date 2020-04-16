package main

import "fmt"

func sendMessage(ch chan string, message string) {
	ch <- message
}

func main() {
	channels := []chan string{}
	messages := []string{
		"A quick brown fox jumps over the lazy dog.",
		"Five hexing wizard bots jump quickly.",
		"How razorback-jumping frogs can level six piqued gymnasts!",
		"Grumpy wizards make toxic brew for the evil queen and jack.",
		"Have a pick: twenty-six letters — no forcing a jumbled quiz!",
	}

	for index := 0; index < len(messages); index++ {
		ch := make(chan string)
		go sendMessage(ch, fmt.Sprintf("%d: %v", index+1, messages[index]))
		channels = append(channels, ch)
	}

	for index := 0; index < len(messages); index++ {
		select {
		case msg := <-channels[index]:
			fmt.Printf("Message received: %v\n", msg)
		}
	}
}
