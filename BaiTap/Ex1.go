//Exercise: Loops and Functions

package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	z := 1.0
	for i := 0; i < 10; i++ {
		z -= (z*z - x) / (2 * z)
		fmt.Printf("Lần lặp thứ %d: z = %v\n", i+1, z)
	}
	return z
}

func main() {

	fmt.Printf("Căn bặc hai của 2: %v\n", Sqrt(2))
	fmt.Printf("Can bậc hai sử dụng hàm sqrt: %v\n", math.Sqrt(2))
}
