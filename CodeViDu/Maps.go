package main

import "fmt"

// type Vertex struct {
// 	Lat, Long float64
// }

// var m map[string]Vertex

// func main() {
// 	m = make(map[string]Vertex)
// 	m["Bell Labs"] = Vertex{
// 		40.68433, -74.39967,
// 	}
// 	fmt.Println(m["Bell Labs"])
// }
///////Một ví dụ khác về map

// type Student struct {
// 	id   int
// 	name string
// 	age  int
// }

// func main() {
// 	//Khoi tao map
// 	studentData := make(map[int]Student)

// 	//them thong tin

// 	studentData[01] = Student{id: 01, name: "Binh", age: 23}
// 	studentData[02] = Student{id: 02, name: "Duong", age: 21}

// 	id := 02
// 	student, exist := studentData[id]
// 	if exist {
// 		fmt.Printf("Id : %d\n", student.id)
// 		fmt.Printf("Name la : %s\n", student.name)
// 		fmt.Printf("Age la : %d\n", student.age)

// 	} else {
// 		fmt.Println("Khong tim thay sinh vien")
// 	}
//}
///////////////////////Map literals
type Vertex struct {
	lat, long float64
}

var m = map[string]Vertex{
	"A": Vertex{40.12345, -70.23123},
	"B": Vertex{12.12345, -80.23123},
}

func main() {
	fmt.Print(m)
}
