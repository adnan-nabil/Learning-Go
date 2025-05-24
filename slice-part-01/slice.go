package main 
import "fmt"

func main(){

	var a [6]string = [6]string{"this", "is", "go", "slice", "important", "topic"}

	// 1. Make slice from array
	var s []string = a[1:5]
	fmt.Println("s = ", s) // pointer: "is" or 1,  length: 3 and capacity: 5
	fmt.Println(cap(s)) // capacity: 5

	// 2. make slice from slice
	var s1[]string = s[1:3]
	fmt.Println("s1 = ", s1) // pointer: "go" or 1, length: 2 and capacity: 4
	fmt.Println(cap(s1)) // capacity: 4

	// 3. make slice by slice literal
	var x[]int = []int{1, 2, 5}
	fmt.Println("x = ", x)
	fmt.Println(cap(x)) 
	fmt.Println(len(x))

	// using make function with length and capacity
	var x1 []int = make([]int, 5, 10) 
	fmt.Println("x1 = ", x1) // [0 0 0 0 0]
	fmt.Println(cap(x1)) // 5
	x1[3] = 115
	fmt.Println("x1 = ", x1) // [0 0 0 115 0]
	x1[7] = 225
	fmt.Println("x1 = ", x1) // [0 0 0 115 0 0 0 225]


	// empty slice or nil slice
	var e []string
	fmt.Println("e = ", e) 

}

// SLICE ER MODDHE KONO ARRAY THAKE NA.
// slice er vitore thake shudhu 3 ta information thake.
// 1. pointer. First value er address
// 2. length. total elemnent	
// 3. capacity. 

//1. slice from array
//2. slice from slice
//3. slice from slice literal
//4. slice from make function with len and capacity
//5. empty slice or nil slice