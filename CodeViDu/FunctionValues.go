package main

import (
	"fmt"
)

// func compute(fn func(float64, float64) float64) float64 {
// 	return fn(3, 4)
// }

// func main() {
// 	hypot := func(x, y float64) float64 {
// 		return math.Sqrt(x*x + y*y)
// 	}

// 	fmt.Println(hypot(5, 12))
// 	fmt.Println(compute(hypot))
// 	fmt.Println(compute(math.Pow))
// }
//////////////////////

func add(a, b int) int {
	return a + b
}

func main() {
	addFunc := add
	result := addFunc(5, 5)
	fmt.Println(result)
}
