package main

import "fmt"

const NewVariable string = "This is a Public variable" // Public variable as it starts with a capital

func main() {
	var name string
	fmt.Printf("hello '%s' world '%s' \n", name, NewVariable)

	var number int
	fmt.Printf("hello '%b' world '%s' \n", number, NewVariable)
}
