package main

import "fmt" // it's bring auto when we use fmt

//jwToken := 3000 // it's not allowed outside the methode

const LoginToken string = "rayene" //LoginToken is public accessible by any file in this folder

func main() {
	var username string = "Amina" //you should use the declaration variable
	fmt.Println(username)         //fp for the shortcut of fmt.Println
	fmt.Printf("Vaiable is of type : %T \n", username)

	var isLogged bool = true
	fmt.Println(isLogged)
	fmt.Printf("Vaiable is of type : %T \n", isLogged)

	var smallVal uint8 = 255 //[0 to 255]
	fmt.Println(smallVal)
	fmt.Printf("Vaiable is of type : %T \n", smallVal)

	///var smallFloat float32 = 255.45875625 // in output i have to get just 5 number after . and the other will ignored
	var smallFloat float64 = 255.45875625 // more precision result
	fmt.Println(smallFloat)
	fmt.Printf("Vaiable is of type : %T \n", smallFloat)

	// default values and some aliases
	var anotherVariable int
	fmt.Println(anotherVariable)
	fmt.Printf("Vaiable is of type : %T \n", anotherVariable)

	//implicit type

	var website = "amina.com" // decided by the lexer
	fmt.Println(website)
	// no var style
	numberOfUser := 3000 //with clon equal
	fmt.Println(numberOfUser)

	fmt.Println(LoginToken)
	fmt.Printf("Vaiable is of type : %T \n", LoginToken)

}
