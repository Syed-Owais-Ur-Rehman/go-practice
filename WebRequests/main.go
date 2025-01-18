package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

const myUrl = "https://www.w3schools.com/"
const myUrl2 = "https://dummy.web:9000/learn?name=js&paymentid=12"

func main() {
	fmt.Println("Handling Web Requests.")

	response, err := http.Get(myUrl)

	if err != nil {
		panic(err)
	}

	fmt.Printf("Response Type is %T \n", response) // will return pointer of orignal response

	defer response.Body.Close()

	databytes, err := ioutil.ReadAll(response.Body) // will scrape the body of webpage
	if err != nil {
		panic(err)
	}

	content := string(databytes)

	fmt.Println("content of Body is:", content)

	fmt.Println("Handling Urls.")
	fmt.Println("New url is: ", myUrl2)

	//parsing
	result, _ := url.Parse(myUrl2)

	fmt.Println("Scheme: ", result.Scheme)
	fmt.Println("Host: ", result.Host)
	fmt.Println("Path: ", result.Path)
	fmt.Println("Port: ", result.Port())
	fmt.Println("Query: ", result.RawQuery) // for all query params

	qparams := result.Query()

	fmt.Printf("Type of Query Params are %T\n", qparams)
	fmt.Println("Value of name in Query Params is: ", qparams["name"])

	for _, val := range qparams {
		fmt.Println("param is: ", val)
	}

	// Create url from parts
	partsOfUrl := &url.URL{
		Scheme:   "https",
		Host:     "dummy.web",
		Path:     "/Admin",
		RawQuery: "user=user1",
	}

	anotherUrl := partsOfUrl.String() // Same as string(partsOfUrl)

	fmt.Println("Another Url: ", anotherUrl)

}
