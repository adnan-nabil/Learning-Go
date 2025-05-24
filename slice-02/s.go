package main
import "fmt"

func main(){
	var x []int
	x = append(x, 1)
	x = append(x, 2)
	x = append(x, 3)

	var y []int = x

	x = append(x, 4)
	y = append(y, 5)

	x[0] = 10

	fmt.Println("x = ", x) // [10, 2, 3, 5]
	fmt.Println("y = ", y) // [10, 2, 3, 5]
}


// simulation na korle bujha jabe na. video te simulation kore bujhano ase. 1.40 theke 1.50 mins er moddhe