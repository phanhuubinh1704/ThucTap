package main

import "fmt"

func main() {
	//////////////////////////Slices
	// a := [5]int{1, 2, 3, 4, 5}

	// var b []int = a[1:5]
	// fmt.Println(b)

	////////////Slices are like references to arrays
	// names := [5]string{"Binh", "Hieu", "Hung", "Tai", "Cuong"}

	// fmt.Println(names)

	// a := names[1:5]
	// b := names[0:1]
	// fmt.Println(a, b)

	// b[0] = "xxx" //thay đổi slices b => thay đổi slices a và mảng names
	// fmt.Println(a, b)
	// fmt.Println(names)

	/////////Slice literals
	// a := []int{1, 2, 3, 4, 5}
	// fmt.Println(a)

	// b := []bool{true, false, true}
	// fmt.Println(b)

	// c := []struct {
	// 	i int
	// 	j bool
	// }{ /// { ko đc xuống dòng
	// 	{1, true}, {2, true}, {3, false},
	// }
	// fmt.Println(c)

	////////////////////Slice defaults
	s := []int{2, 3, 5, 7, 11, 13}

	s = s[1:4]

	fmt.Println(s)

	s = s[:2]
	fmt.Println(s)

	s = s[1:]
	fmt.Println(s)

}
