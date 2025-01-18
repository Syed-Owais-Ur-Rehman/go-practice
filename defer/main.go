package main

import "fmt"

func main() {
	fmt.Println("Defer is like Finally ins Go.")

	// // Will be printed in LIFO order: Hello Two One World
	// defer fmt.Println("World")
	// defer fmt.Println("One")
	// defer fmt.Println("Two")
	// fmt.Println("Hello")

	// Will be printed in LIFO order: Hello 43210 Two One World
	defer fmt.Println("World") // 5
	defer fmt.Println("One")   // 4
	defer fmt.Println("\nTwo") // 3
	fmt.Println("Hello")       // 1
	myDefer()                  // 2

}

func myDefer() {
	for i := 0; i < 5; i++ {
		defer fmt.Print(i)
	}

}
