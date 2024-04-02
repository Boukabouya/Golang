package main

import "fmt"

func main() {
	fmt.Println("alternative version of classes are structs in golang")
	// no inheritence in golang ; Nor super or parent

	gitesh := User{"Gitesh", "Gitesh@go.dev", true, 25} // creating
	fmt.Println(gitesh)
	fmt.Printf("hitesh details are : %+v\n ", gitesh) // %+v for more detail in the struct
	fmt.Printf("the name is %v and emial is %v", gitesh.Name, gitesh.Email)

}

// define structs User with upper case to export it
type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}
