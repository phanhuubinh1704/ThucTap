package main

import "fmt"

func main() {
	var name type [3]string

	a[0] = "Tôi"
	a[1] = "tên"
	a[2] = "Bình"

	//a := [3]string{"Tôi", "tên", "Bình"}

	fmt.Println(a[0], a[1], a[2])
	fmt.Println(a)

	b := [5]int{5, 6, 3, 7, 8}
	fmt.Println(b)
}
