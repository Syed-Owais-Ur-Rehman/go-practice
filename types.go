package main

import "fmt"

func types() {
	var name string
	fmt.Printf("hello '%s' world '%s' \n", name, NewVariable)

	var number int
	fmt.Printf("hello '%b' world '%s' \n", number, NewVariable)
}
