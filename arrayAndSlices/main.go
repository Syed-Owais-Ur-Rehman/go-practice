package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("We are doing arrays and slices")

	// Arrays
	var fruitList [4]string
	fmt.Println("Orignal Fruit List: ", fruitList, "Length of List is: ", len(fruitList))

	fruitList[0] = "Apple"
	fruitList[1] = "Tomato"
	fruitList[3] = "Peach"
	// fruitList[3] = 2 cannot assign different data type

	fmt.Println("Modified Fruit List: ", fruitList, "Length of List is: ", len(fruitList))

	// Slices
	var numberList = []string{"1", "2", "3"}
	fmt.Println("number List Before is:", numberList)

	numberList[0] = "b"

	fmt.Println("number List After is:", numberList)

	// Slices with make()
	highScore := make([]int, 4)

	fmt.Println("highScore Slice: ", highScore)
	fmt.Printf("Item type is: %T \n", highScore[0])

	highScore[0] = 22
	highScore[1] = 33
	highScore[2] = 44
	highScore[3] = 55
	// highScore[4] = 66 Out of Range as Size is 4
	fmt.Println("highScore Slice now is: ", highScore)

	highScore = append(highScore, 88, 66, 77)
	fmt.Println("highScore Slice now is: ", highScore)

	// Sorting
	fmt.Println("Sorted: ", sort.IntsAreSorted(highScore))
	sort.Ints(highScore) // sort imported
	fmt.Println("Sorted: ", sort.IntsAreSorted(highScore))

	// Remove Value from Slice based on Index
	var courses = []string{"react", "js", "python", "swift"}
	var index int = 2

	fmt.Println("courses: ", courses)

	// courses = append(courses[:index], courses[index+1:])
	courses = append(courses[:index], courses[index+1:]...)

	fmt.Println("courses: ", courses)

}
