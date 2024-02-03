package main

import "fmt"

func main() {
	//counter based loop
	for i:=0;i<10;i++ {
		fmt.Println(i)
	}


	//condition loop 
	i :=0
	for  i<10{      
		fmt.Println(i)
		i++
	}

	//infinite loop
	for{
		fmt.Println("Hello world") ///without break this loop will infinite 
		break
	}
}