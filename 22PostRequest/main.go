package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

func main() {
	fmt.Println("The Post request accept the json data")
	//PerformGetRequest()
	PerformPostJsonRequest()
}

func PerformGetRequest() {
	const myurl = "https://www.thunderclient.com/welcome"

	// Create get request
	response, err := http.Get(myurl)
	if err != nil {
		panic(err)
	}

	// close the request
	defer response.Body.Close()

	fmt.Println("Status code:", response.StatusCode)
	fmt.Println("Content length is", response.ContentLength)

	var resposeString strings.Builder // from strings package
	// builder is uded to effeciently build a string using Write methods
	content, _ := io.ReadAll(response.Body)
	byteCount, _ := resposeString.Write(content)

	fmt.Println("ByteCount is ", byteCount)
	fmt.Println(resposeString.String())

	//fmt.Println(content)//bytecode
	//fmt.Println(string(content))
}

// send some data with post request
/*
 to send some data and probably that data might want to store into a database or something
 but one thing is having a surety that some data needs to go in the back end now how you send the data usual cases are in the json format or in the url encoded forms
 in this one let's take down the json
*/
// generate some of the fake data with json

func PerformPostJsonRequest() {
	const myurl = "https://www.thunderclient.com"
	// fake json payload (payload meaning data)

	requestBody := strings.NewReader(`
		// here we could write any kind of data in our case we'll using the json type
		{
			"course":"let's go with golang",
			"price": 0,
			"platform" : "LearnWithAmina.Alg"
		}
	`)
	respons, err := http.Post(myurl, "text/html", requestBody)
	if err != nil {
		panic(err)
	}
	content, _ := io.ReadAll(respons.Body)
	fmt.Println(string(content))

}
