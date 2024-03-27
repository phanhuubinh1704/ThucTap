package main

import "fmt"

////////////////////////////////

// var a, java, golang bool

// func main() {
// 	a, java, golang = true, true, false
// 	var i int = 1
// 	fmt.Print(i, a, java, golang)
// }

// var i, j int = 1, 2

// func main() {
// 	var c, python, java = true, false, "yes"
// 	fmt.Println(i, j, c, python, java)
// }
///////////////////////////////
var k = 3

func main() {
	var i, j int = 1, 2

	c, python, java := true, false, "no"

	fmt.Println(i, j, k, c, python, java)
}
