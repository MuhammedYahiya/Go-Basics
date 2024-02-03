package main

import "fmt"



func main() {
	pointerBasic()
    a := 4
    squareVal(a)
    squareAdd(&a) // after this function, the value of 'a' changes to its square, i.e., from 4 to 16
    fmt.Println(a)
}

// squareVal takes a value 'v' and squares it, printing the result.
func squareVal(v int) {
    v *= v
    fmt.Println("&v:", &v, "v:", v)
}

// squareAdd takes a pointer 'p' to an integer and squares the value it points to,
// effectively changing the original value since it's passed by reference.
func squareAdd(p *int) {
    *p *= *p
    fmt.Println("p:", p, "*p:", *p)
}




func pointerBasic(){
	i,j := 45,27021
	p := &i //storing the address of the variable a so we has the address of i 
	fmt.Println(*p) //*p print the value in at that address	 means value at i
	*p = 21 //change the value on that address means value of i change to 21
	p = &j
	fmt.Println(*p)
}


