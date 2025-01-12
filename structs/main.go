package main

import "fmt"

func main() {
	fmt.Println("Structs are alternate versions of classes.")

	user1 := User{"John Doe", "john.doe@gmail.com", true, 16}

	fmt.Println("User 1 is: ", user1)

	fmt.Printf("User Details are: %v\n", user1)
	fmt.Printf("User Details are: %+v\n", user1)

	fmt.Printf("Name is %v and email is %v\n", user1.Name, user1.Email)

}

type User struct {
	Name   string
	Email  string
	status bool
	Age    int
}
