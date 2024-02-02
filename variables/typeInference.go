package main

import "fmt"

func main() {
	var i int
	j := i // j is an integer
	fmt.Println(j)

	f := 4 + 4.5 //float64
	g := 0.865 + 0.5i //complex128
	fmt.Println(f,g)
}