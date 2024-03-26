package main

import (
	"fmt"
	"math"
)

// func sqrt(x float64) string {
// 	if x < 0 {
// 		return sqrt(-x) + "i"
// 	}
// 	return fmt.Sprint(math.Sqrt(x))
// }
// func main() {
// 	fmt.Println(sqrt(2), sqrt(-4))
// }

func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	}
	return lim
}

func main() {
	fmt.Println(pow(1, 2, 9), pow(4, 3, 15))
}
