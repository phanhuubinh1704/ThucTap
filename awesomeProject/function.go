package main

import "fmt"

func Chao() {
	fmt.Println("Hello")
}

func Chao1(name string) {
	fmt.Println("Hello", name)
}
func greeting(name string) string {
	result := fmt.Sprint("Hello %s", name)
	return result

}
func main() {
	Chao()
	Chao1("Binh")
	result := greeting("A")
	fmt.Println(result)
}
