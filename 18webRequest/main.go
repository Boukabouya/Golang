package main

import (
	"fmt"
	"io"
	"net/http"
)

const url = "https://httpbin.com/"

func main() {
	// you should close connection of the response of request after use it
	fmt.Println("welcome to web request")

	response, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Response id of type : %T\n", response)

	defer response.Body.Close() // always close body after using, caller's responsibility to close the connection

	// read the response
	databytes, err := io.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}
	// convert the data byte to the content
	content := string(databytes)
	fmt.Println(content)

}

//Response id of type : *http.Response
// we get a reference of all original response that we have recieved
// so the pointer is garantee that we are not recieve any copy of response and it's the
