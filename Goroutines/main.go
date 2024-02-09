package main

import (
	"fmt"
	"time"
)

// greeter function prints the provided string 6 times with a slight delay
func greeter(s string) {
	// Loop 6 times
	for i := 0; i < 6; i++ {
		// Introduce a delay of 3 milliseconds
		time.Sleep(3 * time.Millisecond)
		// Print the provided string
		fmt.Println(s)
	}
}

func main() {
	// Start a new goroutine to execute the greeter function with "hello" as argument
	go greeter("hello")
	// Call the greeter function with "world" as argument in the main goroutine
	greeter("world")
}
