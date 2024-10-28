package main

import "fmt"

// variable := "hello"  := type can only be declared inside functions
func main() {

	// fmt.Println(variable)  //this doesnt work
	var a = "initial"
	fmt.Println(a)

	var b, c int = 1, 2
	fmt.Println(b, c)

	var d = true
	fmt.Println(d)

	var e int
	fmt.Println(e)

	f := "apple"
	fmt.Println(f)



	var x, y, z int = 10, 20, 30 // declare multiple variables
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)
	var (
		i int    = 420
		j int    = 69
		k string = "puck yea"
	)

	fmt.Println(i)
	fmt.Println(j)
	fmt.Println(k)

	var p, q = 6, "hello" // types get detected
	fmt.Println(p)
	fmt.Println(q)

	m, n := true, false
	fmt.Println(m)
	fmt.Println(n)


	myVariableName := "John"  //camel case
	fmt.Println(myVariableName)

	MyVariableName := "John"   // pascal case
	fmt.Println(MyVariableName)

	my_variable_name := "John"  // snake case
	fmt.Println(my_variable_name)
}
