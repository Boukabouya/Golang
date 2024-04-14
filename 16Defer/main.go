package main

import "fmt"

func main() {
	fmt.Println("Defer Act by LIFO, it's run by reverse ")

	defer fmt.Println("world") //last out
	defer fmt.Println("one")   // 3rd out
	defer fmt.Println("two")   // 2nd out

	fmt.Println("hello") // it's run by reverse //1 st out
	fmt.Println("it's me ")
	myDefer()
}
// 0,1,2,3,4
// hello,it's me,43210,two,one,world 

func myDefer() {
	for i := 0; i < 5; i++ {
		defer fmt.Println(i)
	}
}
