package main

import (
	"encoding/json"
	"fmt"
)

// type course struct {
// 	Name     string
// 	Price    int
// 	Platform string
// 	Password string
// 	Tags     []string
// }

type course struct {
	Name     string `json:"coursename"`
	Price    int
	Platform string   `json:"website"`
	Password string   `json:"-"`
	Tags     []string `json:"tags,omitempty"`
}

func main() {
	fmt.Println("Creating and Consuming Json Data.")

	// Creating Json Data
	EncodeJson()

	// Consuming Json Data
	DecodeJson()

}

func EncodeJson() {
	dummyCourse := []course{
		{"React JS", 299, "dummy.web1.com", "abc123", []string{"web-dev", "js"}},
		{"MERN", 499, "dummy.web2.com", "def123", []string{"full-stack", "js"}},
		{"Angular", 184, "dummy.web3.com", "xyz123", nil},
	}

	// Package this data as Json data
	// finalJson, err := json.Marshal(dummyCourse)
	finalJson, err := json.MarshalIndent(dummyCourse, "", "\t")

	if err != nil {
		panic(err)
	}

	fmt.Println(string(finalJson))

}

func DecodeJson() {
	jsonDataFromWeb := []byte(`
		{
			"coursename": "React JS",
			"Price": 299,
			"website": "dummy.web.com",
			"tags": ["web-dev","js"]
		}
	`)

	var dummyCourse course

	checkValid := json.Valid(jsonDataFromWeb)

	if checkValid {
		fmt.Println("Json is valid")

		json.Unmarshal(jsonDataFromWeb, &dummyCourse)
		fmt.Printf("%#v\n", dummyCourse)
	} else {
		fmt.Println("Json was not valid")
	}

	// Some case where you just want to add data to key value pair
	var myOnlineData map[string]interface{} // interface used here coz we don't know input data.

	json.Unmarshal(jsonDataFromWeb, &myOnlineData)

	fmt.Printf("%#v\n", myOnlineData)

	for k, v := range myOnlineData {
		fmt.Printf("Key: %v Value: %v Type: %T\n", k, v, v)
	}

}
