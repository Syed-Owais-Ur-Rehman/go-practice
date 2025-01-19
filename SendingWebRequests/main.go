package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

func main() {
	fmt.Println("There are 3 to perform web requests: Get, Post with Json data, and Psot with data in urlencoded format.")

	// GET
	PerformGetRequests()

	// Post with JSON data
	PerformPostJsonRequests()

	// Post with JSON Form
	PerformPostFormRequests()
}

func PerformGetRequests() {
	const myUrl = "http://localhost:8000/get" // assuming a frontend is runnig that returns that we provide in body.

	response, err := http.Get(myUrl) // run only if dummy server is up

	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	fmt.Println("Status Code: ", response.StatusCode)
	fmt.Println("Content Length: ", response.ContentLength)

	content, _ := ioutil.ReadAll(response.Body)
	fmt.Println(string(content)) // Although easy, the following method is more usually used.

	// We can replace the above string(content) with following preferred manner:
	var responseString strings.Builder // Builder module used for more functionalities
	byteCount, _ := responseString.Write(content)
	fmt.Println("Byte Count is: ", byteCount)
	fmt.Println(responseString.String()) // responseString has now more functions to work with

}

func PerformPostJsonRequests() {
	const myUrl = "http://localhost:8000/post" // assuming a frontend is runnig that returns that we provide in body.

	requestBody := strings.NewReader(` // dummy Json data returned when sent to frontend.
		{
			"coursename": "ReactJS",
			"price": 0
		}
	`)

	response, err := http.Post(myUrl, "application/json", requestBody)

	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	content, _ := ioutil.ReadAll(response.Body)
	fmt.Println(string(content))

}

func PerformPostFormRequests() {
	const myUrl = "http://localhost:8000/postform" // assuming a frontend is runnig that returns that we provide in body.

	// Form Data
	data := url.Values{}
	data.Add("firstname", "john")
	data.Add("lastname", "doe")
	data.Add("age", "16")

	response, err := http.PostForm(myUrl, data)

	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	content, _ := ioutil.ReadAll(response.Body)
	fmt.Println(string(content))
}
