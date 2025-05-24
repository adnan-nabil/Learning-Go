package main

import "fmt"

var a = 20

func main(){
	fmt.Println("ami main function er vitore")
	fmt.Println(a)
}

func init(){
	fmt.Println("ami age print hobo pore main")
	a = 30
}