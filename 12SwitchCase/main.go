package main

import (
	"fmt"
	"math/rand"
)

func main() {
	fmt.Println("Switch and case in golang")
	/*You can remove rand.Seed(time.Now().UnixNano())
	from your Go 1.20+ code for generating random numbers
	 each time you run the program.
	rand.Seed(time.Now().UnixNano())// random number every single time
	*/
	diceNumer := rand.Intn(6) + 1 // cuz 6 mean betwween 0 and 5 so i add 1 to skip the 0
	fmt.Println("Value of dice is ", diceNumer)

	// get start with switch
	switch diceNumer {
	case 1:
		fmt.Println("Dice value is 1 and you can open")
	case 2:
		fmt.Println("You can move to 2 spot")
	case 3:
		fmt.Println("You can move to 3 spot")
		fallthrough // mean run the next case also
	case 4:
		fmt.Println("You can move to 4 spot")
		fallthrough
	case 5:
		fmt.Println("You can move to 5 spot")
	case 6:
		fmt.Println("You can move to 6 spot and roll dice again")
	default:
		fmt.Println("what was that ?")
	}
}
