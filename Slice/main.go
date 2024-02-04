package main

import (
	"fmt"
	"sort"
)

func main() {

	//initializing slice
	fruits := []string{"Apple", "Orange", "Banana"}

	//Accessing value in slice
	fmt.Println(fruits)
	fmt.Println(fruits[0])
	fmt.Println(fruits[1])
	fmt.Println(fruits[2])

	//added new element to slice
	fruits = append(fruits, "Mango", "Avocado")
	fmt.Println(fruits)

	//we can take only some parts of the slice using

	fruits = fruits[0:4] //it will avoid 4 the element
	fmt.Println(fruits)

	//Removing elements from slice based on index
	name := []string{"Yahiya", "Akshay", "Suhail", "Aswin", "Abi"}
	index := 2
	name = append(name[:index], name[index+1:]...) //this will remove 2nd index element from slice
	fmt.Println(name)

	//We can slicing from an array
	prime := [6]int{2, 3, 5, 7, 11, 13}
	s := prime[1:4]
	fmt.Println(s)

	//slice are like reference to array
	country := [5]string{"India", "Australia", "New Zealand", "England", "South Africa"}

	a := country[0:2]
	b := country[1:3]
	fmt.Println(a, b)
	//I wanted to change the country Australia to Austria
	b[0] = "Austria" //this will change to the country array
	fmt.Println(country)

	//Creating slice with make

	highScores := make([]int, 4)
	highScores[0] = 234
	highScores[1] = 947
	highScores[2] = 465
	highScores[3] = 123
	//highScores[4] = 233 // this will show out of bound error
	highScores = append(highScores, 555, 666, 321) //this will added the values to highScores
	fmt.Println(highScores)

	//sorting
	sort.Ints(highScores) //sorting the integer values
	fmt.Println(highScores)
	//check the elements are sorted
	fmt.Println(sort.IntsAreSorted(highScores))

}
