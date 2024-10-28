package main

import "fmt"

func main() {
	fmt.Println("Value from values.go!")

	var i string = "Hello"
	var j int = 15

	fmt.Printf("i has value: %v and type: %T\n", i, i)
	fmt.Printf("j has value: %v and type: %T", j, j)
}
