package main

import (
	"fmt"
	"math"
)

// Define the shape interface, which requires an area method.
type shape interface {
	area() float64
}

// Define the rectangle struct with width and height fields.
type rectangle struct {
	width  float64
	height float64
}

// Define the circle struct with a radius field.
type circle struct {
	radius float64
}

// Implement the area method for the rectangle struct.
func (r rectangle) area() float64 {
	return r.height * r.width
}

// Implement the area method for the circle struct.
// Note: This method has a pointer receiver to modify the original circle instance.
func (c *circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

// Function to calculate the area of any shape by calling its area method.
func getArea(s shape) {
	fmt.Println(s.area())
}

func main() {
	// Create a rectangle instance with width 7 and height 5.
	r1 := rectangle{7, 5}
	// Create a circle instance with a radius of 4.5.
	c1 := circle{4.5}

	// Create a slice of shape interface containing the rectangle and circle instances.
	shapes := []shape{r1, &c1}

	// Iterate over the shapes slice and print the area of each shape.
	for _, shape := range shapes {
		fmt.Println(shape.area())
	}
	getArea(r1)
}
