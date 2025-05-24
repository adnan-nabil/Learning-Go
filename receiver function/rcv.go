package main

import "fmt"

type Person struct{
	name string
	age int
	sex string
}

// receiver function
func (p Person) printdetails(){
	fmt.Println("Name: ", p.name)
	fmt.Println("Age: ", p.age)
	fmt.Println("sex: ", p.sex)
}

func main(){

	nahid := Person{
		name: "Nahid",
		age: 25,
		sex: "male",
	}

	nahid.printdetails()

	var nabil Person = Person{
		name: "Nabil",
		age: 25,
		sex: "male",
	}

	nabil.printdetails()
}

func init(){
	fmt.Println("this is a receiver function init code")
}