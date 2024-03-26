package main

import "fmt"

// type Vertex struct {
// 	x int
// 	y int
// }

// func main() {
// 	fmt.Println(Vertex{1, 2})
// }

//////Struct Fields

// type Vertex struct {
// 	X int
// 	Y int
// }

// func main() {
// 	v := Vertex{1, 2}
// 	v.X = 4 //gán giá trị mới
// 	fmt.Println(v.X)
// }

/////Pointers to structs
type Vertex struct {
	x int
	y int
}

func main() {
	v := Vertex{4, 5}
	p := &v
	p.x = 1e7
	fmt.Println(v)
}
