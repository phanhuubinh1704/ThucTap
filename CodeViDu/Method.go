package main
//Method cũng là một Function nhưng method thuộc về một đối tượng cụ thể
///khai báo method func (t Type) methodName(params) returns{//body code}
///(t type ) gọi là Receiver
//Có 2 loại Receiver
//1 Value Receiver
//2 Pointer Receiver
import (
	"fmt"
	"math"
)

type Vertex struct {
	x, y float64
}

func (v Vertex) Abs() float64 { // method Abs thuộc về đối tượng Vertex
	return math.Sqrt(v.x*v.x + v.y*v.y)
}
//////Một ví dụ khác 
type Student struct {
	name string
}

func (s Student) getName() string {
	return s.name
}

func main() {
	v := Vertex{3, 4}
	fmt.Println(v.Abs())
	fmt.Println("===========================")
	student := Student{"Binh"}
	name := student.getName()
	fmt.Println(name)
}
//Có 2 loại Receiver
//1 Value Receiver khi thay đổi giá trị thay đổi giá trị trong func ko ảnh hường đền func main 
//2 Pointer Receiver
//(t *type)

///////////////////////Methods and pointer indirection
type Vertex struct {
	X, Y float64
}

func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func ScaleFunc(v *Vertex, f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func main() {
	v := Vertex{3, 4}
	v.Scale(2)
	ScaleFunc(&v, 10)

	p := &Vertex{4, 3}
	p.Scale(3)
	ScaleFunc(p, 8)

	fmt.Println(v, p)
}
















