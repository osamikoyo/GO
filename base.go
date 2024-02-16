package main

import "fmt"
import "math"

func sp(count float64, step float64){
	  resusd:= math.Pow(count, step)
	  fmt.Println(resusd)
	
}
func ab(x int, y int){
	fmt.Println("x befor = ", x,  "y befor = ", y)
	x = x + 20; y = y + 20
	fmt.Println("x after = " ,x, "y after = ", y)
}

func main() {
	sp(10, 5)
	ab(4, 6)
	}


