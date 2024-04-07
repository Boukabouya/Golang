package main

import "fmt"

// main func act as an entrypoint of programme without call it
func main() {
	fmt.Println("welcome to functions")
	greeter()

	greeterTow()

	result := adder(3, 5)
	fmt.Println("Result is :", result)
	proREs, myMessage := proAdder(2, 5, 7, 9)
	fmt.Println("ProResult is :", proREs)
	fmt.Println("my message is :", myMessage)

}
func greeterTow() {
	fmt.Println("Another method")
}
func greeter() {
	fmt.Println("marhaba from golang")
}

// int at the end is a signiture of the return value
func adder(valOne int, valTow int) int {
	return valOne + valTow
}

// now i want to declare func without knowing the input parameter
func proAdder(values ...int) (int, string) {
	//value are slice
	total := 0
	for _, val := range values {
		total += val
	}
	return total, "hi from Pro Adder function"
}
