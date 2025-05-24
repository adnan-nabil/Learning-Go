package main

import "fmt"

// standard or named function. Jar nam ase
func add(a int, b int) int {
	res := a+b
	return res
}

func main(){
	ans := add(5,7)

	fmt.Println((ans))

	// immidiately invoked function expression, jar kono nam nai. tai nogod nogod call diye dite hobe.
	// eijonnoi nam immidiately invoked function 
	func(x int, y int){
		fmt.Println(x+y)
	}(10, 15)
}

func init(){
	fmt.Println(" Ami shobar age call hoye jabo")
}