package main

import (
	Test "GoLang/calling/parentCall.go"
	"fmt"
)

func main(){
	fmt.Println("am inside ./parentFolder/file")
	Test.Testing("Hello am cuming from parent")
}