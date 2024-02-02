package main

import "fmt"

func main() {
	var num1, num2 int = 1,2 //initialized variable
	var name, age, drink = "Yahiya", 22, false // here we emitted the type because the variable will take type of initialization
	k := 3 // := short assignment statement can be used in place of var declaration with implicit type declaration
	fmt.Println(num1,num2,name,age,drink,k)

}