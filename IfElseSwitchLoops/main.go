package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("If Else, Switch and Loops")

	// If-Else
	loginCount := 23
	var result string

	if loginCount < 10 {
		result = "Regular User"
	} else if loginCount > 10 {
		result = "Watch Out"
	} else {
		result = "Exactly 10"
	}

	fmt.Println("Result is: ", result)

	if 9%2 != 0 {
		fmt.Println("This is another type of If Else")
	}

	if num := 3; num < 10 {
		fmt.Println("This is another type of If Else where a variable is declared beforehand")
	}

	// Error handling with if
	// if err != nil {

	// }

	// Switch

	// Create a new random source using the current time as the seed
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)

	// rand.New allows each goroutine to have its own independent random
	// rand.Seed sets the global seed for the shared random number generator, which can lead to issues in concurrent programs.

	// Generate random numbers using the new Rand instance
	fmt.Println(r.Intn(100)) // Random integer between 0 and 99
	fmt.Println(r.Float64()) // Random float between 0.0 and 1.0

	diceNum := r.Intn(6) + 1

	fmt.Println("Dice Num is: ", diceNum)

	switch diceNum {

	case 1:
		fmt.Println("Dice value is 1")
		fallthrough

	case 2:
		fmt.Println("Dice value is 2")

	case 3:
		fmt.Println("Dice value is 3")

	case 4:
		fmt.Println("Dice value is 4")
		fallthrough // falls to one case below and exits

	case 5:
		fmt.Println("Dice value is 5")

	case 6:
		fmt.Println("Dice value is 6")

	default:
		fmt.Println("What was that")

	}

	// Loops
	days := []string{"Sun", "Tues", "Wed", "Fri", "Sat"}

	for d := 0; d < len(days); d++ { // This loop won't be used much
		fmt.Println("Day:", days[d])
	}

	for i := range days { // easier loop method
		fmt.Println("Day is:", days[i]) // i here does not have value but has index
	}

	for index, day := range days {
		fmt.Printf("index is %v and value is %v \n", index, day)
	}

	for _, day := range days {
		fmt.Printf("index is and value is %v \n", day)
	}

	// Equivalent to while loop

	rogueValue := 1

	// Style 1 : Normal
	for rogueValue < 10 {
		fmt.Println("Value: ", rogueValue)
		rogueValue++
	}

	rogueValue = 1

	// Style 2 : Using Break
	for rogueValue < 10 {

		if rogueValue == 5 {
			break
		}

		fmt.Println("Loop with Beak, Value: ", rogueValue)
		rogueValue++
	}

	rogueValue = 1

	// Style 3 : Using Continue without Increment
	for rogueValue < 10 {

		if rogueValue == 5 {
			rogueValue++ // <----- if this is commented then value of rogueValue as 5 will continue to be traversed infinitely.
			continue
		}

		fmt.Println("Loop with Continue and without increment, Value: ", rogueValue)
		rogueValue++

	}

	rogueValue = 1

	// Style 4 : Using Continue with Increment
	for rogueValue < 10 {

		if rogueValue == 5 {
			rogueValue++
			continue
		}

		fmt.Println("Loop with Continue and Increment, Value: ", rogueValue)
		rogueValue++

	}

	rogueValue = 1

	// Style 5 : Using label with loop
	for rogueValue < 10 {

		if rogueValue == 2 {
			goto lco
		}

		fmt.Println("Loop with Label, Value: ", rogueValue)
		rogueValue++

	}

lco:
	fmt.Println("Jumping to Code")
}
