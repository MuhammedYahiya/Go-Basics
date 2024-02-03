package main

import "fmt"

func main() {

	//creating an instance of the User struct
	user1 := User{
		Name:   "Yahiya",
		Email:  "yahiya@gmail.com",
		Status: true,
		Age:    22,
	}
	//Accessing fields of the User struct
	fmt.Println(user1)
	fmt.Println("Name:", user1.Name)
	fmt.Printf("User1 details are: %+v\n", user1)                   //User1 details are: {Name:Yahiya Email:yahiya@gmail.com Status:true Age:22}
	fmt.Printf("Name is %v and Age is %v\n", user1.Name, user1.Age) //Name is Yahiya and Age is 22

	// Creating instance of car struct
	car := []Car{
		{
			Name:  "Honda",
			color: "Black",
		},
		{
			Name:  "Toyota",
			color: "White",
		},
		{
			Name:  "BMW",
			color: "Blue",
		},
	}

	// Accessing fields
	fmt.Println(car)
	fmt.Println("Name:", car[0].Name)
	fmt.Println("Color:", car[0].color)
	fmt.Println("Name:", car[1].Name)
	fmt.Println("Color:", car[1].color)
	fmt.Println("Name:", car[2].Name)
	fmt.Println("Color:", car[2].color)
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

type Car struct {
	Name  string
	color string
}
