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
	//EncodeJson()
	DecodeJson()
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

// consume  or decode the JSON data which is converted from struct

func DecodeJson() {
	// any data come from the web it's actually in the byte format we need to wrap it into string
	// but here the byte data we'll consume it as a json not convert it into string
	jsonDataFromWeb := []byte(`
	{
		"coursename": "ReactJS Bootcamp",
		"Price": 299,
		"website": "LearnCodeOnline.in",
		"tags": ["web-dev","js"]
	}
	
	`)
	//to verify if my data  is valid Json or not I use json.Valid function
	// it's a commun practice in golang
	var lcoCourse course

	checkValid := json.Valid(jsonDataFromWeb) // return boolean so it' doesn't give us an err
	if checkValid {
		fmt.Println("This is Valid Json")
		// now we can Unmarshall our Json data into our Struct
		// & means pass by reference so it will change the value of our variable directly
		// if we don't put & then it will create new copy of varibale and changes won't reflect outside
		// because Go follow Copy Elimination concept
		json.Unmarshal(jsonDataFromWeb, &lcoCourse)
		fmt.Printf("%#v\n", lcoCourse) // %# special syntax for print the interfaces

	} else {
		fmt.Println("This is Invalid Json")
	}

	//some cases where you just want to add data to key value
	// key is guarantee that it's a sting however the value  could be anything
	var myOnlineData map[string]interface{}
	json.Unmarshal(jsonDataFromWeb, &myOnlineData)
	fmt.Printf("%#v\n", myOnlineData) // by print it it's look complex
	// so by make it in loop it's look simple
	for k, v := range myOnlineData {
		fmt.Printf("key is %v and value is %v and type is %T \n", k, v, v)
	}

}
