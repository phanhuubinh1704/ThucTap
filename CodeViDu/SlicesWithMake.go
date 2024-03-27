package main

import "fmt"

func main() {
	//Hàm make cấp phát một mảng chứa các phần tử được khởi tạo
	a := make([]int, 7)
	printSlice("a", a)

	b := make([]int, 0, 7)
	printSlice("b", b)

	c := b[:3]
	printSlice("c", c)

	d := c[2:6]
	printSlice("d", d)
}
func printSlice(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n", s, len(x), cap(x), x)
}
