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
// variables

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
}
