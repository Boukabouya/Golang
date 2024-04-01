package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("welcome to the slice")
	// slice syntax
	var fuitsList = []string{}                          // by this syntax you should inisialise it
	fmt.Printf("Type of fruitslist is %T\n", fuitsList) // type of the fruitList

	var fuitsList2 = []string{"Apple", "Tomato", "Peach"} // inisialise it by value
	fmt.Printf("Type of fruitslist is %T\n", fuitsList2)  // type of the fruitList2

	//Add values
	fuitsList2 = append(fuitsList2, "Mongo", "Banana") //  the first arg of append we add the slice that we want to append it after we add the value
	fmt.Println(fuitsList2)
	// the most append used usally
	fuitsList2 = append(fuitsList2[1:3]) // we slice our slice stop at position 3 but it's non inclusive
	// we use it alot exemple in the todolist app or anything related to the database
	fmt.Println(fuitsList2)
	/*In summary, the first approach creates an empty slice with no initial memory allocation for elements,
	  while the second approach creates a slice with a specific capacity
	  and allocates memory upfront.
	*/
	highScores := make([]int, 4) // we allocate 4 space for our slice
	//by default this is an array getting a layer of abstraction and thus we call them slices so technically they are slices
	highScores[0] = 234
	highScores[1] = 235
	highScores[2] = 236
	highScores[3] = 237
	//highScores[4]=777 //out of the bounds
	highScores = append(highScores, 555, 354, 458) //append will reallocate the memory and the new value will be accommodated

	fmt.Println(highScores)
	fmt.Println(sort.IntsAreSorted(highScores)) // false //we test if the slice is sorted or not
	//  available with slice not in the array to sort value we specified that it's int data type
	sort.Ints(highScores)
	fmt.Println(highScores)

	// how to remove a value from slice based on the index
	var courses = []string{"reactJs", "javascript", "python", "ruby", "swift"}
	fmt.Println(courses)
	var index int = 2
	// append coude use it to remove a value
	// super important
	courses = append(courses[:index], courses[index+1:]...)
	fmt.Println(courses)

}
