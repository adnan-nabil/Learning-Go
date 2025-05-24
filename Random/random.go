package main

import "fmt"

func main(){

	var(
		p = 10
		q = 20
	)

	add := func(x int, y int){
		result := x + y
		fmt.Println(result)
	}

	add(p, q)

}