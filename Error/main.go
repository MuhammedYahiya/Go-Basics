package main

import (
	"errors"
	"fmt"
)

// divide function takes two float64 numbers and returns the result of division
// It returns an error if the divisor is zero
func divide(x, y float64) (float64, error) {
	// Check if the divisor is zero
	if y == 0 {
		// If divisor is zero, return an error indicating division by zero
		return 0, errors.New("division by zero")
	}
	// If divisor is not zero, perform division and return the result
	return x / y, nil
}

func main() {
	// Call divide function with parameters 10 and 0
	result, err := divide(10, 0)
	// Check if there was an error returned from the divide function
	if err != nil {
		// If there was an error, print the error message
		fmt.Println("Error:", err)
	} else {
		// If there was no error, print the result of the division
		fmt.Println("Result:", result)
	}
}
