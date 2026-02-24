package main

import (
	"fmt"
)

func main() {
	messages := make(chan string, 3)

	fmt.Println("Sending messages to buffered channel")
	messages <- "First Message"
	messages <- "Second message"
	messages <- "Third message"
	// messages <- "Third message" doesn't work because is set on 3

	fmt.Println(<-messages)
	fmt.Println(<-messages)
	fmt.Println(<-messages)

}
