package main

import (
	"fmt"
	"math"
)

func main() {
	var x, y int = 3, 4
	var f float64 = math.Sqrt(float64(x*x + y*y))
	var z uint = uint(f)
	fmt.Println go run Type.go(x, y, z)

	v := 0.867 + 0.5i
	b := 1.332
	fmt.Printf("v is of type %T\n", v)
	fmt.Printf("b is of type %T\n", b)
}
