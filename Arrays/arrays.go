package main
import "fmt"

func main(){
	var a[6] int
	fmt.Println("empty array:", a)


	a[4] = 5
	fmt.Println("set", a)
	fmt.Println("get", a[4])

	fmt.Println("length of array a is:", len(a))

	b:= [5] int{1,2,3,4,5}
	fmt.Println("array b is:", b)

	b = [...]int{100, 200, 300, 400, 500}
    fmt.Println("Redeclaring array b :", b)



	var x,y,z int = 10,20,30;
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)
}