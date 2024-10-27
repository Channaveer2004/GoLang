package main

import (
	"fmt"
	"time"
)

func main() {
	i := 2
	fmt.Println("Write", i, "as")
	switch i {
	case 1:
		fmt.Println("one")

	case 2:
		fmt.Println("two")

	case 3:
		fmt.Println("three")
	}

	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("its weekend bish")

	default:
		fmt.Println("weedays sugs")
	}


		t:= time.Now()
		switch  {
		case t.Hour() < 12:
			fmt.Println("morning mf")
			
			default :
			fmt.Println("afternoon !")
		}

		whatami := func(i interface{}){
			switch t := i.(type){
			case bool:
				fmt.Println("Am Bool")
			case int:
				fmt.Println("Am integer")
			case string:
				fmt.Println("Am string")
			default:
				fmt.Println("Ifdk wt type",t)
			}
		}

		whatami(true)
		whatami(69)
		whatami("true")
		// whatami(func(whatami))
		

}
