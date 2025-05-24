package main

import "fmt"

func main(){

	var result int = 100
	var d int = 50
	fmt.Println("Hello, World!")

	// function expression. Anonymous function. a function in a variable like  python
	// I have to call it after the declaration inside a function. but if i declare it globally its okay to call anywhere
	f := func(x int, y int){
		result = x+y
		fmt.Println(d)
	}

	f(2, 3)

	fmt.Println(result)
}

func init(){
	fmt.Println("Init function")
}