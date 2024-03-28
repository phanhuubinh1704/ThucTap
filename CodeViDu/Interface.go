package main

import (
	"fmt"
	"math"
)

type Abser interface {
	Abs() float64
}

type MyFloat float64

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

type Vertex struct {
	X, Y float64
}

func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// //////////////////Nil interface values
type I interface {
	M()
}

func describe(i I) {
	fmt.Printf("(%v,%T)\n", i, i)
}

// /////////////The empty interface
func describe1(j interface{}) {
	fmt.Printf("(%v,%T)\n", j, j)
}
func main() {
	var a Abser
	f := MyFloat(-math.Sqrt2)
	v := Vertex{3, 4}

	a = f
	a = &v
	fmt.Println(a.Abs())

	fmt.Println("==============Nil interface values=================")
	var j I
	describe1(j) // In ra: "(<nil>,<nil>)"
	// j.M() //
	fmt.Println("================The empty interface=================")
	var k interface{}
	describe1(k)

	k = 42
	describe1(k)

	k = "hello"
	describe1(k)
}
