package main

import "fmt"

func main() {
	fmt.Println("welcome to loops")
	days := []string{"sunday", "Tuesday", "wednesday", "Friday", "Satruday"}

	fmt.Println(days)
	for d := 0; d < len(days); d++ {
		fmt.Printf("%v\n", days[d])
	}
	// for without initialization
	for i := range days {
		fmt.Println(days[i])
	}
	for index, day := range days {
		fmt.Printf("%v for the index %v\n", day, index)
	}
	rougValue := 1
	for rougValue < 10 {
		if rougValue == 2 {
			goto lco
		}
		if rougValue == 5 {
			break // stop the for loop and out
		}
		if rougValue == 6 {
			rougValue++
			continue //skip and continue so if we don't increament the value it's crach the execusion
		}

		fmt.Println(rougValue)
		rougValue++
	}
lco: //lco like a pointer
	fmt.Println("Jumping at LearnAminaCode.in")

}
