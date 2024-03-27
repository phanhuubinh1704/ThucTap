package main

import "fmt"

func main() {
	var a []int
	printSlice(a)
	//Hàm append thêm phần tử vào slices
	a = append(a, 0)
	printSlice(a)

	a = append(a, 1)
	printSlice(a)

	a = append(a, 1, 2, 3)
	printSlice(a)

}

func printSlice(a []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(a), cap(a), a)
}
