package main

import "fmt"

func main() {

	user1 := User{
		Name:   "Muhammed Yahiya",
		Email:  "yahiya@gmail.com",
		Age:    22,
		Status: true,
	}
	fmt.Println(user1)
	user1.GetStatus()
	user1.NewEmail()         // this function only make a copy
	fmt.Println(user1.Email) // output is yahiya@gmail.com

}

type User struct {
	Name   string
	Email  string
	Age    int
	Status bool
}

func (u User) GetStatus() {
	fmt.Println("User status is: ", u.Status)
}

func (u User) NewEmail() {
	u.Email = "test@gmail.com"             //this only make the copy not change in the main user
	fmt.Println("New email is: ", u.Email) //test@gmail.com
}
