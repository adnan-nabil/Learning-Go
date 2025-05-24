package main

import "fmt"

func math(num1 int , num2 int) (int,int, string) {
	sum := num1+num2
	mul := num1*num2
	s := "ami nabil"
	return sum, mul, s
}

func nameprint(name string){
	fmt.Println("my name is",name,".I am 26 years old")
}


func main() {
	sum, mul, name := math(2,5)
	fmt.Println(sum, mul, name)
	nameprint("nabil")	

}