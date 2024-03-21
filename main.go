package main

/*import (
	"fmt"
	"math"
	"math/rand"
)

func main() {
	fmt.Println("Hello world")
	fmt.Println(rand.Int())
	fmt.Printf("The squar is %g that's all \n", math.Sqrt(9))
}*/
// variables exemple 01
/*
import (
	"fmt"
	"math/cmplx"
)

func main() {
	var i int = 1
	var f float32 = 2.75
	var f2 float64 = 2.758846
	var s string = "Hello go"
	var b bool = true
	var c complex128 = cmplx.Sqrt(5 + 12i)
	fmt.Println(i, f, f2, s, b, c)
}*/
// Variables exemple 02
/*import (
	"fmt"
)

func main() {
	var i int
	//i = 5
	//i = 27
	//fmt.Println("\n The value  of i is ", i)//The value  of i is  27
	var f float32
	var s string
	var b bool
	fmt.Printf("%v %v %q %v\n", i, f, s, b) //%v : default print , %q  : for string with quote
}

// Type Inference
import "fmt"

func main() {
	var i int
	i = 1 

	var j int = 5
	k := 9 //Type inference: Go automatically determines the type based on the context of the variable
	var (
		a string = "hello"
		b string = "amina"
		c int    = 25
	)
	fmt.Println(i, j, k)
	fmt.Println(a, b, c)
}

//Reassigning Variables
import "fmt"
func main(){
	var i int = 4
	i = 45 //overwriting the i value
	//i := 20 // don't allow to  overwrite the i value because it has been declared using 
	fmt.Println("New Value of I is ", i)
}
*/

// Variable scope
import "fmt"

var i int = 4 
func main(){
	var i int = 5 // scope variable , more prioritaire than the outside main func  var i 
	fmt.Println(i)
}
