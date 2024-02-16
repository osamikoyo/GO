package main

import (
	"fmt"
	"math"
)

func sp(count float64, step float64) {
	resusd := math.Pow(count, step)
	fmt.Println(resusd)

}
func ab(x int, y int) {
	fmt.Println("x befor = ", x, "y befor = ", y)
	x = x + 20
	y = y + 20
	fmt.Println("x after = ", x, "y after = ", y)
}
func add(x, y int) int {
	fmt.Println(x, " + ", y, " = ", x+y)
	return x + y
}


var mun int = 10
var p *int
var k = &mun
func main() {
	sp(10, 5)
	ab(4, 6)
	add(10, 657)
	var gul int = add(5,7)
	fmt.Println(gul)
	fmt.Println(*k)
	n := new(int)
	fmt.Println("value = ", *n)

	*n = 87
	fmt.Println("value = ", *n)
}
