package main

import "fmt"

func add(x int, y int, z int) {
	sum := x + y + z
	fmt.Println(sum)
}

func add2(x, y, z int) int { // here added int because we returning 2 values sum
	return x + y + z
}

func swap(x, y string) (string, string){ // here returned two string not doing any other operation just return 2 string so we use (string,string)
	return y,x
}

func main() {
    add(1,2,3)
	add2(3,4,5)
	x,y := swap("hello", "world")
	fmt.Println(x,y)
}