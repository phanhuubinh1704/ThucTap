package main

import (
	"fmt"
	"math"
)

func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		fmt.Printf("%g <= %g\n", v, lim)
		return v
	} else {
		fmt.Printf("%g >= %g\n", v, lim)
		return lim
	}
}

func main() {
	fmt.Println(
		pow(4, 3, 20),
		pow(2, 3, 30),
	)
}
