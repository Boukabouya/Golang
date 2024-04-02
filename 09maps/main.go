package main

import "fmt"

/*
format key value
the retrieval are very fast in maps
easaly delete the values as well
*/
func main() {

	fmt.Println("welcome to the maps")

	//create a maps
	languages := make(map[string]string) //make is allowed for maps also and maps here have a string for the key and string for vslue data type
	// add value
	languages["JS"] = "Javascript"
	languages["RB"] = "Ruby"
	languages["PY"] = "Python"
	fmt.Println("List of all the languages", languages)
	fmt.Println("JS shorts for: ", languages["JS"])

	// delete the value from maps
	delete(languages, "RB")
	fmt.Println("List of all the languages", languages)

	// loops are interesting in golang
	for key, value := range languages {
		fmt.Printf("For key %v, value is %v\n", key, value)
	}

	// if we don't care about the key we just replace it with underscore to ignore it
}
