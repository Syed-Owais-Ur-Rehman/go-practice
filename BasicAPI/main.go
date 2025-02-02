package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

type Course struct {
	CourseId    string  `json:"courseid"`
	CourseName  string  `json:"coursname"`
	CoursePrice int     `json:"price"`
	Author      *Author `json:"author"`
}

type Author struct {
	Fullname string `json:"fullname"`
	Website  string `json:"website"`
}

// fake DB
var courses []Course

func main() {
	fmt.Println("Building a basic API in Go.")

	r := mux.NewRouter()

	// Seeding
	courses = append(courses, Course{
		CourseId:    "2",
		CourseName:  "React JS",
		CoursePrice: 120,
		Author: &Author{
			Fullname: "John",
			Website:  "abc.com",
		}})
	courses = append(courses, Course{
		CourseId:    "1",
		CourseName:  "Angular",
		CoursePrice: 240,
		Author: &Author{
			Fullname: "Doe",
			Website:  "abc2.com",
		}})

	// Routing
	r.HandleFunc("/", serveHome).Methods("GET")
	r.HandleFunc("/courses", getAllCourses).Methods("GET")
	r.HandleFunc("/course/{id}", getOneCourse).Methods("GET")
	r.HandleFunc("/course", createOneCourse).Methods("POST")
	r.HandleFunc("/course/{id}", updateOneCourse).Methods("PUT")
	r.HandleFunc("/course/{id}", deleteOneCourse).Methods("DELETE")

	// Listen to a port
	log.Fatal(http.ListenAndServe(":4000", r))
}

// Method of Course type
// Middleware, helper - File --> separate file
func (c *Course) IsEmpty() bool {
	return c.CourseId == "" && c.CourseName == ""
}

// controller - File

// Serve Home Route
func serveHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Welcome to a basic API</h1>"))
}

func getAllCourses(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get All Courses")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(courses)
}

func getOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get One Course")
	w.Header().Set("Content-Type", "application/json")

	// grab id from request
	params := mux.Vars(r)

	// loop through courses, find matching id and return the response
	for _, course := range courses {
		if course.CourseId == params["id"] { // id in params
			json.NewEncoder(w).Encode(course)
			// return // redundant statement but used

		}

	}
	json.NewEncoder(w).Encode("No Course found with given id")
	// return // redundant statement but used

}

func createOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Create One Course")
	w.Header().Set("Content-Type", "application/json")

	// what if: body is empty
	if r.Body == nil {
		json.NewEncoder(w).Encode("Please send some data")
	}

	// data sent like {}
	var course Course

	_ = json.NewDecoder(r.Body).Decode(&course)

	if course.IsEmpty() {
		json.NewEncoder(w).Encode("No data inside JSON")
		// return
	}

	// generate unique id, string
	// append course into courses

	// rand.Seed(time.Now().UnixNano())  // Deprecated
	rand.New(rand.NewSource(time.Now().UnixNano()))
	course.CourseId = strconv.Itoa(rand.Intn(100)) // rand.Intn(100) used for case where id is taken as integer.

	courses = append(courses, course)
	json.NewEncoder(w).Encode(course)
	//return

}

func updateOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Create One Course")
	w.Header().Set("Content-Type", "application/json")

	// First - grab id from request
	params := mux.Vars(r)

	var courseFound bool = false

	// loop through courses, find matching id , remove and add with my ID
	for index, course := range courses {
		if course.CourseId == params["id"] { // id in params

			courseFound = true

			courses = append(courses[:index], courses[index+1:]...)

			var course2 Course
			_ = json.NewDecoder(r.Body).Decode(&course2) // pushing request body params into course2
			course2.CourseId = params["id"]
			courses = append(courses, course2)
			json.NewEncoder(w).Encode(course2)
			// return // redundant statement but used
			break

		}

	}

	if !courseFound {
		json.NewEncoder(w).Encode("No Course found with given id")
		// return // redundant statement but used
	}

}

func deleteOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Delete One Course")
	w.Header().Set("Content-Type", "application/json")

	// First - grab id from request
	params := mux.Vars(r)

	var courseFound bool = false

	// loop, id , remove (index, index+1)
	for index, course := range courses {
		if course.CourseId == params["id"] { // id in params

			courseFound = true

			courses = append(courses[:index], courses[index+1:]...)

			json.NewEncoder(w).Encode("Course Item deleted") // TODO task

			break

		}

	}

	if !courseFound {
		json.NewEncoder(w).Encode("No Course found with given id")
		// return // redundant statement but used
	}
}
