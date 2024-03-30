package main

import "fmt"

func main() {
	fmt.Println("welcome to the pointer")
	//Pointers exist to ensure that the actual value of a variable is passed on when variables are passed to functions.
	//Create a pointer
	//Pointers are created using the asterisk (*) symbol and are used to reference memory addresses.
	var ptr *int // holding integers into that
	// the default value of pointer
	fmt.Println("the default value of the pointer is ", ptr) //nil

	// create a variable that assigned to another variable
	myNumber := 23
	var ptr1 = &myNumber                    //& mean refrencing
	fmt.Println("values of the ptr", ptr1)  // output = 0xc000012090 cus it's refrence to the memory location
	fmt.Println("values of the ptr", *ptr1) // output = 23 *ptr1 mean i want to see what is inside the ptr
	*ptr1 = *ptr1 + 2
	fmt.Println("New value is : ", myNumber) // 25 Pointers can be used to perform operations on the actual value of a variable, rather than on copies of the value.
	/*
		Pointers are useful in cases where variables need to be passed to multiple functions.
	*/

}
