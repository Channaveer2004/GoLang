package main

import (
	"fmt"
	"math"
)

const s string = "constant"      // can be declared outside functions also
const Apple_1 = "apple_1"        // typed constant
const Apple_2 string = "apple_2" // untyped constant

func main() {
	fmt.Println(s)
	fmt.Println(Apple_1)
	fmt.Println(Apple_2)

	const n = 500000000
	const d = 3e20 / n
	fmt.Println(d)

	fmt.Println(int64(d))

	fmt.Println(math.Sin(n))

	const cantChange = "cant change const value once declared"
	fmt.Println(cantChange)

	const (
		A int = 1
		B     = 3.14
		C     = "Hi!"
	)

	fmt.Println(A)
	fmt.Println(B)
	fmt.Println(C)
}
