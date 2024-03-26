package main

import "fmt"

type Vertex struct {
	x int
	y int
}

// func main() {
// 	fmt.Println(Vertex{1, 2})
// }

//////Struct Fields

// func main() {
// 	v := Vertex{1, 2}
// 	v.X = 4 //gán giá trị mới
// 	fmt.Println(v.X)
// }

/////Pointers to structs

// func main() {
// 	v := Vertex{4, 5}
// 	p := &v
// 	p.x = 1e7
// 	fmt.Println(v)
// }
///////Struct Literals

var (
	v1 = Vertex{2, 2}
	v2 = Vertex{x: 3}
	v3 = Vertex{}
	p  = &Vertex{1, 1}
)

func main() {
	fmt.Println(v1, p, v2, v3)
}
