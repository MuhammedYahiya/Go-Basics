package main

import "fmt"

var python,java bool //declaration without initialization outside the main function	

func main() {
	var num int;  //declaration without initialization inside the main function	
	var name string
	fmt.Println(python,java,num,name)
	// the output should be false false 
	// the default value of boolean is false int is zero string is empty string

}