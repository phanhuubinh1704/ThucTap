package main

import "fmt"

func sl(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

func main() {
	fmt.Print(sl(20))
}
