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

// build complete API without data base by gorila mux package by handling the routing

// Model for course - file
type Course struct {
	CourseId    string  `json:"courseid"`
	CourseName  string  `json:"coursename"`
	CoursePrice int     `json:"price"`
	Author      *Author `json:"author"` //custom type so type is a pointer
}

type Author struct {
	Fullname string `json:"fullname"`
	Website  string `json:"website"`
}

// fake DB
var courses []Course

// middleware, helper - file
// if we get an empty courseid and name we declare that the field is empty
func (c *Course) IsEmpty() bool {
	// return c.CourseId == "" && c.CourseName == ""
	// cuz the course id willl generate manually
	return c.CourseName == ""
}

func main() {
	fmt.Println("API - LearnCodeOnline.in")
	r := mux.NewRouter()

	//seeding
	courses = append(courses, Course{CourseId: "2", CourseName: "ReactJS", CoursePrice: 299, Author: &Author{Fullname: "Hitesh Choudhary", Website: "lco.dev"}})
	courses = append(courses, Course{CourseId: "4", CourseName: "MERN Stack", CoursePrice: 199, Author: &Author{Fullname: "Hitesh Choudhary", Website: "go.dev"}})

	//routing
	r.HandleFunc("/", serveHome).Methods("GET")
	r.HandleFunc("/courses", getAllCourses).Methods("GET")
	r.HandleFunc("/course/{id}", getOneCourse).Methods("GET")
	r.HandleFunc("/course", createOneCourse).Methods("POST")
	r.HandleFunc("/course/{id}", updateOneCourse).Methods("PUT")
	r.HandleFunc("/course/{id}", deleteOneCourse).Methods("DELETE")

	// listen to a port
	log.Fatal(http.ListenAndServe(":4000", r))
}

//controllers - file

// serve home route
// r -> get the request from somebody else , w -> write the answer : header or ...
func serveHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Welcome to API by LearnCodeOnline</h1>"))
}

func getAllCourses(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get all courses")
	//set a header
	w.Header().Set("Content-Type", "applicatioan/json")
	// return the response
	//write data in json format and back to the one how request the data
	json.NewEncoder(w).Encode(courses)
}

func getOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get one course")
	w.Header().Set("Content-Type", "applicatioan/json")

	// grab id from request
	/*Extracting Variable with mux.Vars(r):
	return a map contains key-value pairs
	mux.Vars(r) acts like a dictionary to retrieve values captured from URL variables in your Gorilla Mux routes.
	*/
	params := mux.Vars(r)

	// loop through courses, find matching id and return the response
	for _, course := range courses {
		if course.CourseId == params["id"] {
			json.NewEncoder(w).Encode(course)
			return
		}
	}
	json.NewEncoder(w).Encode("No Course found with given id")
	return
}

func createOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Create one course")
	w.Header().Set("Content-Type", "applicatioan/json")

	// what if: body is empty
	if r.Body == nil {
		json.NewEncoder(w).Encode("Please send some data")
	}

	// what about - {}

	var course Course
	_ = json.NewDecoder(r.Body).Decode(&course)
	if course.IsEmpty() {
		json.NewEncoder(w).Encode("No data inside JSON")
		return
	}

	//TODO: check only if title is duplicate

	// loop, title matches with course.coursename, JSON

	// generate unique id, string
	// append course into courses

	rand.Seed(time.Now().UnixNano())               // we don't need it it's stoped with go 1.20
	course.CourseId = strconv.Itoa(rand.Intn(100)) // rand.Intn(100) give me an integer
	courses = append(courses, course)
	json.NewEncoder(w).Encode(course)
	return

}

func updateOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Update one course")
	w.Header().Set("Content-Type", "applicatioan/json")

	// first - grab id from req
	params := mux.Vars(r)

	// loop, id, remove, add with my ID

	for index, course := range courses {
		if course.CourseId == params["id"] {
			courses = append(courses[:index], courses[index+1:]...) // remove course  by index
			var course Course
			_ = json.NewDecoder(r.Body).Decode(&course)
			course.CourseId = params["id"]
			courses = append(courses, course)
			json.NewEncoder(w).Encode(course)
			return
		}
	}
	//TODO: send a response when id is not found
	// responce not found or something like that
	// err := errors.New("Course with the given ID was not found.")
	err := fmt.Sprintf("Course with the given ID %v was not found.", params["id"])
	http.Error(w, err, http.StatusNotFound)
	return
}

func deleteOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Delete one course")
	w.Header().Set("Content-Type", "applicatioan/json")

	params := mux.Vars(r)

	//loop, id, remove (index, index+1)

	for index, course := range courses {
		if course.CourseId == params["id"] {
			courses = append(courses[:index], courses[index+1:]...)
			//break
			// TODO: send a confirm or deny response
			json.NewEncoder(w).Encode(course) // return deleted course
			return
		}
	}
	// if we are here it means there's no such course in our system
	// so let's return an error message to inform about this issue
	err := fmt.Sprintf("Course with the provided ID %s does not exist.", params["id"])
	http.Error(w, err, http.StatusNotFound)
}
