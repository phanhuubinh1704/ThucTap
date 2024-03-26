package main

import "fmt"

// func main() {
// 	defer fmt.Println("A")

// 	fmt.Println("B")
// }

//Stacking defers
func main() {
	fmt.Println("counting")

	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}

	fmt.Println("done")
}
