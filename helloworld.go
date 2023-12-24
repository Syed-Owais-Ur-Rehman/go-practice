package main

import (
	"fmt"
)

// := only used inside function block
var b float32 = 3.4
var c, d int = 1, 7

func main() {

	var (
		f1   float64
		int1 int = 89
		b1   bool
		s1   = "string"
	)

	var str1 string = "richard"
	var str2 string = "lionheart"
	var e, f, g = "3", 4, false
	h, i := 17.5, "value!"

	/*

		Camel case
		myVariableName = "John"

		Pascal case
		MyVariableName = "John"

		Snake case
		my_variable_name = "John"

	*/

	x := "3"
	y := true
	z := -16

	fmt.Println("Here are some strings: ", str1, str2, x)
	fmt.Println("some other values: ", y, z, b)
	fmt.Println("some global variable values: ", c, d)
	fmt.Println("some local variable values: ", e, f, g)
	fmt.Println("some other: ", h, i)
	fmt.Println("var block vals: ", f1, int1, b1, s1)

	const PI, A = 3.14, 6
	const B, C int = 3, 4

	a2 := 2

	fmt.Println("Constant Values", a2, PI, A, B, C)
}
