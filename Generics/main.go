package main

import "fmt"

// add function uses generics to accept arguments of any numeric type (int or float64)
func add[T int | float64](num1 T, num2 T) T {
	return num1 + num2 // Adds two numbers and returns the result
}

func main() {
	sum := add(1.2, 2) // Calls the add function with float64 and int arguments
	fmt.Println(sum)
}
