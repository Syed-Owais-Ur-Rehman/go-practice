package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Println("We'll also explore Files in Go.")

	// Creating and Writing in file.
	content := "This needs to go in file."

	file, err := os.Create("./myText.txt")

	NilErr(err)
	// if err != nil {
	// 	panic(err)
	// }

	length, err := io.WriteString(file, content)

	NilErr(err)
	// if err != nil {
	// 	panic(err)
	// }

	fmt.Println("length is: ", length)

	defer file.Close()

	// Reading File
	readFile("./myText.txt")
}

func readFile(filename string) {
	databyte, err := ioutil.ReadFile(filename)

	NilErr(err)
	// if err != nil {
	// 	panic(err)
	// }

	fmt.Println("Text data in file:", string(databyte)) // works like parseFloat & parseInt
}

func NilErr(err error) {
	if err != nil {
		panic(err)
	}
}
