package main

import "fmt"

func main() {
	var i int = 48
	var f float64 = float64(i)
	var g uint = uint(f)
	fmt.Println(i,f,g)
}