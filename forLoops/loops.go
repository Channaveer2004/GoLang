package main
import "fmt"


func main() {
	fmt.Println("hey sup");

	var i int = 1;
	for i<=3{
		fmt.Println(i)
		i++
	}

	for j:=10; j<=15; j++{
		fmt.Println(j)
	}

	for i:= range 3{
		fmt.Println("Range", i)
	}

	for{
		fmt.Println("Break")
		break
	}

	for n:=10; n<20;n++{
		if(n%2==0){
			continue
		}
		fmt.Println(n)
	}

}