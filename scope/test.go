package main

import "fmt"


var (
	a = 10
	b = 20
)

func printNum(res int){
	fmt.Println(res)
}

func addNum(a int, b int){
	res := a+b
	printNum(res)
}

func main(){
	addNum(a, b)
}