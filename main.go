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
import (
	"fmt"
)

func main() {
	var i int
	/*i = 5
	i = 27
	fmt.Println("\n The value  of i is ", i)//The value  of i is  27*/
	var f float32
	var s string
	var b bool
	fmt.Printf("%v %v %q %v\n", i, f, s, b) //%v : default print , %q  : for string with quote
}
