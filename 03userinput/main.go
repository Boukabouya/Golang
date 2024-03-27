package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "welcome to user program"
	fmt.Println(welcome)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the rating of our pizza :")
	// comma ok syntax || err ok : to store the reader input in variable
	// cuz the golang don't use the try catch
	/*
		Error handling in Go involves using the comma ok
		or comma error syntax to capture errors and ignore values.
		 This syntax allows for efficient error management in Go programming.
		 and it's like we treat a varibles or true and false value
	*/
	//use _ in the place of err or input if you don't have to care about it
	input, _ := reader.ReadString('\n') // store the reader in input var
	//so if i use the err like i catch the try
	fmt.Println("Thanks for reading, ", input)
	fmt.Printf(" Type of this rating is %T ", input)

}
