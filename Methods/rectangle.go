package main // Declares that this is a main package, which can be executed as a standalone program.

import "fmt"

type rect struct { // Defines a struct type named "rect" to represent a rectangle.
	width  int // Field representing the width of the rectangle.
	height int // Field representing the height of the rectangle.
}

func (r rect) perim() int { // Defines a method named "perim" with a value receiver "r" of type "rect".
	// Calculates and returns the perimeter of the rectangle using the value receiver.
	return 2*r.width + 2*r.height
}

func (r *rect) area() int { // Defines a method named "area" with a pointer receiver "r" of type "*rect".
	// Calculates and returns the area of the rectangle using the pointer receiver.
	return r.width * r.height
}

func main() { // The entry point of the program.
	r := rect{width: 10, height: 5} // Creates a rectangle instance with width 10 and height 5.

	// Prints the area and perimeter of the rectangle "r" directly using value receiver methods.
	fmt.Println("Area: ", r.area())
	fmt.Println("Perimeter: ", r.perim())

	rp := &r // Creates a pointer to the rectangle instance "r".

	// Prints the area and perimeter of the rectangle "r" accessed via a pointer using pointer receiver methods.
	fmt.Println("area: ", rp.area())
	fmt.Println("Perimeter: ", rp.perim())
}
