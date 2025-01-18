package main

import "fmt"

func main() {
	fmt.Println("Methods in Go.")

	user1 := User{"John", "john.doe@gmail.com", true, 16}

	user1.GetStatus()

	// It is a deep copy/orignal value does not change
	fmt.Println("Email before change:", user1.Email)
	user1.NewEmail()
	fmt.Println("Email after change:", user1.Email) // Therefore, Pointers used.

}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

func (u User) GetStatus() {
	fmt.Println("is the user active:", u.Status)
}

func (u User) NewEmail() {
	u.Email = "test@gmail.com"
	fmt.Println("Email is now:", u.Email)
}
