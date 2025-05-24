package main

import "fmt"

type Person struct{
	name string
	age int
}

// pass by value, i mean copy
func show(b [3]int){
	fmt.Println("Inside show function")
	fmt.Println(b)
}

// pass by reference, i mean pointer
func newf(number *[3]int){
	fmt.Println(number)
}

// pointer of a struct in a receiver function
func (x *Person) show2(){
	fmt.Println(x)
}

func main(){

	x := 10
	p := &x 
	fmt.Println("Value of x:", x) // 10
	fmt.Println("Address of x:", &x) // Address of x
	fmt.Println("Pointer p:", p) // Address of x
	fmt.Println("Value of p:", *p) // 10

	*p = 20 // Change value of x using pointer p
	fmt.Println("New value of x:", x) // 20

	var a [3]int = [3]int{1, 2, 3}
	show(a)
	newf(&a)



	var nabil Person = Person{
		name: "Nabil",
		age: 25,
	}

	nabil.show2()
}