package main

import "fmt"

func main() {
	fmt.Println("Functions in Go")

	result := adder(3, 5)
	result2 := proAdder(3, 5, 8, 9, 10, 13)
	proRes, myMessage := proAdder2(3, 5)

	fmt.Println("Result 1 is: ", result)
	fmt.Println("Result 2 is: ", result2)
	fmt.Printf("Result 3 is: %v and secret message is: %v ", proRes, myMessage)
}

func adder(valOne int, valTwo int) int {
	return valOne + valTwo
}

func proAdder(values ...int) int { // ... are variadic functions and can expect any value.
	total := 0

	for _, val := range values {
		total += val
	}

	return total
}

func proAdder2(val1 int, val2 int) (int, string) {

	result := val1 + val2

	return result, "Hello"
}
