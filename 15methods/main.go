package main

import "fmt"

func main() {
	// since we have not a classes we have a struct and we need to bring our function into struct
	// we called them methods
	fmt.Println("when functions go into classes we call them method")
	// no inheritence in golang ; No super or parent

	gitesh := User{"Gitesh", "Gitesh@go.dev", true, 25} // creating
	fmt.Println(gitesh)
	fmt.Printf("hitesh details are : %+v\n ", gitesh) // %+v for more detail in the struct
	fmt.Printf("the name is %v and emial is %v\n", gitesh.Name, gitesh.Email)

	//use the methose
	gitesh.GetStatus()
	gitesh.NewMail()
	fmt.Printf("the name is %v and emial is %v\n", gitesh.Name, gitesh.Email)

}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
	//oneAge int // not exportable
}

func (u User) GetStatus() {
	fmt.Println("Is user active:", u.Status)
}

func (u User) NewMail() {
	u.Email = "test@go.dev"
	fmt.Println("New Email of this user is :", u.Email)
}
