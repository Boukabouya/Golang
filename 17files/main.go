package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	// we need os package and  defer
	fmt.Println("welcome to files in golang")
	content := "this need to go in a file"
	file, err := os.Create("./mylcogofile.txt")

	if err != nil {
		panic(err) //shut down the programm and show me the error that i am facing
	}

	length, err := io.WriteString(file, content)

	checkNilErr(err)
	fmt.Println("length is :", length)
	defer file.Close() // the recommanded way is to add defer
	readFile("./mylcogofile.txt")
}

func readFile(filename string) {
	databyte, err := os.ReadFile(filename)

	checkNilErr(err)
	fmt.Println("Text data inside the file is \n", databyte)
	fmt.Println("Text data inside the file is \n", string(databyte)) // simple way to wrap and around to a string

}
func checkNilErr(err error) {
	if err != nil {
		panic(err)
	}
}
