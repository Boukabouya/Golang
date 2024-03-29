package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("welcome to time study of glang")

	//it's useful but not with tith format
	// nanosecond is a commun practive and it give a precise value
	// presentTime := time.Now().Nanosecond() // but with nano we could not use format
	presentTime := time.Now()
	fmt.Println(presentTime)

	//change format and ti=his is the standard  way to do it in GoLangfrom the documentation
	fmt.Println(presentTime.Format("01-02-2006 15:04:05 Monday"))

	createdDate := time.Date(2020, time.August, 10, 23, 23, 0, 0, time.UTC)
	fmt.Println("The date was created on :", createdDate)
	fmt.Println(createdDate.Format("01-02-2006 15:04:05 Monday"))

}
