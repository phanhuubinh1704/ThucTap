package main

import "fmt"

func swapp(x, y string) (string, string) {
	return y, x
}

func main() {
	a, b := swapp("hello", "Binh")
	fmt.Println(a, b)
}
