package main

import "fmt"

func main() {

	fmt.Println("welcome to the array")
	var fruitList [4]string
	/*
		Arrays in Golang are declared using square brackets
		and the specific data type and length must be specified.
		The length of an array is fixed and does not change based
		 on the number of values assigned to it.
	*/
	fruitList[0] = "Apple"
	fruitList[1] = "Tomato"
	fruitList[3] = "peach"
	fmt.Println("Fruit list is :", fruitList)
	//Fruit list is : [Apple Tomato  peach] there are a long space between tomato and peach cus there are item that is not inisialised
	fmt.Println("Fruit list is :", len(fruitList)) //it's strange thing in golang cuz it's run the same number we declared however the inisialisation

	var vegList = [3]string{"potato", "beans", "mushroom"}
	fmt.Println("Vegy List is ", len(vegList))

}
