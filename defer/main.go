package main

import "fmt"

func main() {

	defer fmt.Println("world") // this will only print last
	fmt.Println("hello") 

	fmt.Println("start")
	defer cleanup() // The function call will be deferred until the surrounding function (main) finishes
	fmt.Println("End")
	defer clear()
	defer fmt.Println("yup")
}

func cleanup(){
	fmt.Println("Clean")
}
func clear(){
	fmt.Println("clear")
}

// The order of the output is 
/*
hello
start
end
yup
clear 
clean
world
*/