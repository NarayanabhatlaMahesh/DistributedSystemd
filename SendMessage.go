package main

import (
	"fmt"
	"time"
)

func main() {
	// Create a channel to communicate between goroutines
	messageChannel := make(chan string)

	// Start the first goroutine that sends messages
	go func() {
		for i := 0; i < 5; i++ {
			message := fmt.Sprintf("Message %d from goroutine 1", i)
			messageChannel <- message
			time.Sleep(500 * time.Millisecond) // Simulate some work
		}
		close(messageChannel) // Close the channel when done sending
	}()

	// Start the second goroutine that receives messages
	go func() {
		for message := range messageChannel {
			fmt.Println("Received:", message)
		}
	}()

	// Sleep to allow goroutines to finish
	time.Sleep(3 * time.Second)
}
