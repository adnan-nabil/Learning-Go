package main
import "fmt"

var data  [5]string

func main(){

	data = [5]string{"a", "b", "c", "d", "e"}
	fmt.Println(data[0])

	// way 01
	var s [3]float64 = [3]float64{1.1, 2.2, 3.3}
	fmt.Println(s[2])

	// way 02
	var d [3]string
	d = [3]string{"a", "b", "c"}
	fmt.Println(d[1])

	// way 03 (inside any function)
	x := [3]int{1, 2, 3}
	fmt.Println(x[0])
}