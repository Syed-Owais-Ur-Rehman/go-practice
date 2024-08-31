package main

import "fmt"

func main() {

	// var anyVariableName = [length] type { values }

	var arr1 = [2]int{1, 2}
	myEmptyArr1 := [4]string{}
	myEmptyArr2 := [4]int{}
	arr2 := [5]string{"A", "B"}
	arr3 := [5]int{10, 15}

	// index 3 is out of bounds (>= 3)
	// arr4 := [3]string{"A", "B", "C", "D"}

	arr5 := [...]int{1, 2, 3, 4, 5}
	// Assigning values to index
	arr6 := [5]int{1: 10, 4: 40}

	var names = [3]string{"A", "B", "C"}
	names[1] = "test"

	fmt.Println(arr1)
	fmt.Println(myEmptyArr1)
	fmt.Println(myEmptyArr2)
	fmt.Println(arr2)
	fmt.Println(arr3)
	// fmt.Println(arr4)
	fmt.Println(arr5)
	fmt.Println(len(arr5))
	fmt.Println(names)
	fmt.Println(arr6)

	mySlice := []string{"AB", "BC", "CD"}

	fmt.Println(mySlice)
	fmt.Println(len(mySlice))
	fmt.Println(cap(mySlice)) // Checking Capacity

	arr7 := []string{"AB", "BC", "CD", "XY", "YZ", "ZZ"}

	mySlice2 := arr7[2:4]

	fmt.Println(mySlice2)
	fmt.Println(len(mySlice2))
	fmt.Println(cap(mySlice2)) // Capacity of arr7 will be added

	mySlice3 := make([]int, 5, 10) // if capacity not defined then it will be the same as length defined.

	fmt.Println(mySlice3)
	fmt.Println(len(mySlice3))
	fmt.Println(cap(mySlice3))

}
