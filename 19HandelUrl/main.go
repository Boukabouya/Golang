package main

import (
	"fmt"
	"net/url"
)

const myurl string = "https://www.example.com:3000/learn?googleguy=googley&paimentid=gfhfgfd5356"

func main() {
	// the halper way id net library it's actually a modual that inclue some pices of code
	fmt.Println("handeling urls in golang")
	fmt.Println(myurl)

	//parsing : extract information in our meaning
	result, _ := url.Parse(myurl)
	fmt.Println(result.Scheme)
	fmt.Println(result.Host)
	fmt.Println(result.Path)
	fmt.Println(result.Port())
	fmt.Println(result.RawQuery)

	qparams := result.Query()
	fmt.Printf("The type of query params are : %T\n", qparams)

	fmt.Println(qparams["googleguy"]) // to extraction key of of query that want to parse
	fmt.Println(qparams["paimentid"]) // to extraction key of of query that want to parse

	// loops for all the query
	for _, val := range qparams {
		fmt.Println("Params is ", val)
	}
	// the reverse direction
	// always you pass on refrence not on copy
	partsOfUrl := &url.URL{
		Scheme:  "https",
		Host:    "www.example.com",
		Path:    "/learn",
		RawPath: "user=amina",
	}
	anotherURL := partsOfUrl.String()
	fmt.Println(anotherURL)
}
