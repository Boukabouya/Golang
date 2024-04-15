package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

func main() {
	fmt.Println("Welcome to the Get request in golang")

	PerformGetRequest()
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

/*
In the provided code snippet, the `resposeString.Write(content)` method call writes the content of the `content` variable directly to the `resposeString` builder's buffer. The `bytecode` variable is not involved in this process; it is simply used to store the number of bytes written by the `Write` method. Here's a breakdown:

1. `content, _ := io.ReadAll(response.Body)`: Reads the response body into the `content` variable.

2. `byteCount, _ := resposeString.Write(content)`: Writes the content of the `content` variable to the `resposeString` builder's buffer. The `Write` method returns the number of bytes written, which is stored in the `byteCount` variable.

3. `fmt.Println("ByteCount is ", byteCount)`: Prints the number of bytes written to the builder's buffer.

4. `fmt.Println(resposeString.String())`: Prints the final string built by the builder, which includes the content written in step 2.

Therefore, even though the `bytecode` variable is defined, it is not used to add content to the buffer. Instead, the `content` variable's content is directly added to the `resposeString` builder's buffer through the `Write` method.



*/
