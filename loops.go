package main

import "fmt"

func main() {

	x := 5
	y := 6

	if x > y {
		fmt.Println("x greater than  y")
	} else if x == y {
		fmt.Println("both equal")

	} else {
		fmt.Println("y greater than x")
	}

	//////// Switch Case ////////

	day := 4

	// switch day {
	// case 1:
	// 	fmt.Println("M")
	// case 2:
	// 	fmt.Println("T")
	// case 3:
	// 	fmt.Println("W")
	// case 4:
	// 	fmt.Println("TH")
	// case 5:
	// 	fmt.Println("F")
	// default:
	// 	fmt.Println("Invalid")
	// }

	switch day {
	case 1, 2, 3:
		fmt.Println("First")
	case 4, 5, 6:
		fmt.Println("Second")
	default:
		fmt.Println("Invalid")
	}

	//////// For Loop ////////

	for i := 0; i < 5; i++ {
		if i == 1 {
			continue
		}

		if i == 3 {
			break
		}

		fmt.Println(i)
	}

	names := [3]string{"A", "B", "C"}

	for ids, val := range names {
		fmt.Printf("%v\t %v \n", ids, val)
	}

}
