package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Welcometo our pizza app")
	fmt.Println("Please rate our pizza between 1 and 5")

	reader := bufio.NewReader(os.Stdin) //the new reader get read from os.stdin
	input, _ := reader.ReadString('\n') /// keep to read untel you \n
	fmt.Println("Thanks for reading", input)
	// in sum reason somtimes we want to add 1 to the rating that it's input
	numRating, err := strconv.ParseFloat(strings.TrimSpace(input), 64)
	//strconv.ParseFloat: parsing "5\n": invalid syntax with (input, 64)
	if err != nil { // the err not nil and there are somthing inside it
		fmt.Println(err)
	} else {
		fmt.Println("Added 1 to your rating", numRating+1)
	}
}
