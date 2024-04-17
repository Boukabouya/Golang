package main

import (
	"encoding/json"
	"fmt"
)

// create structure
type course struct {
	Name     string `json:"coursename"` //field of the structure , json tag is used to specify a different name as aliases to look in our API
	Price    int
	Platform string   `json:"website"`        // json tag is used to customize the key name in JSON data
	Password string   `json:"-"`              // The field with a "-" will not be exported,whoever is consuming my api
	Tags     []string `json:"tags,omitempty"` // If Tags has no value ,it won't show up in JSON data.
}

// in this main we'll encoded the struct data it's mean to convert data into valide by json
func main() {
	fmt.Println("Create data with json in golang ")
	EncodeJson()
}

func EncodeJson() {
	// slice of course
	lcoCourses := []course{
		{"ReactJS Bootcamp", 299, "LearnCodeOnline.in", "abc123", []string{"web-dev", "js"}},
		{"MERN Bootcamp", 199, "LearnCodeOnline.in", "bcd123", []string{"full-stack", "js"}},
		{"Angular Bootcamp ", 90, "LearnCodeOnline.in", "hit123", nil},
	}
	//package this data as JSON data for consumable
	//finalJson, err := json.Marshal(lcoCourses) //output a byte slice
	finalJson, err := json.MarshalIndent(lcoCourses, "", "\t") //output a byte slice// for the empty "" in each tap it print the prefix inside ""

	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", finalJson)

}
