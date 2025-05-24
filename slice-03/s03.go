package main
import "fmt"

func changeslice(s []int) []int {
	s[0] = 33
	s = append(s, 11)
	return s
}

func main(){
	var x []int = []int{1, 2, 3, 4, 5}
	x = append(x, 6)
	x = append(x, 7)

	var a []int = x[4:]

	y := changeslice(a)

	fmt.Println("x = ", x)
	fmt.Println("y = ", y)

}