package main

import "fmt"

func main() {

	//array declaration
	var array[2] int
	fmt.Println(array)

	//array initialization
	number := [5]int{1, 2, 3, 4,5}
	fmt.Println(number)

	array[0] = 10 //insert value on array in index 0
	array[1] = 20 //insert value on array in index 1
	fmt.Println(array[0]) //select element at index 0 means position 1

	//length of the array
	length := len(number) //length\
	fmt.Println(length)
}