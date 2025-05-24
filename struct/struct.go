package main

import "fmt"

func init(){
	fmt.Println("Struct shikhar class...")
}

type Student struct{  // exported struct. because of capital S
	name string
	age int
	reg_number int
	address string
}

type teacher struct{ // unexported struct. because of small t
	name string
	subject string
}

func main(){
	fmt.Println("testing 1 2 3 4....")

	nafiz := Student{  // instance of Student
		name : "nafiz",
		age : 14,
		reg_number : 2018331553,
		address: "demra, dhaka",
	}

	harun := teacher{
		name : "harun molla",
		subject : "bengali",
	}


	fmt.Printf("Name: %s\nAge: %d\nReg No: %d\nAddress: %s\n", 
	nafiz.name, nafiz.age, nafiz.reg_number, nafiz.address)


	fmt.Printf("Name = %s\nSubject = %s", harun.name, harun.subject)
}