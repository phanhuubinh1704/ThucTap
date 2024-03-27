package main

import "fmt"

func main() {
	var a = []int{1, 2, 4, 8, 16, 32, 64, 128}
	//Cú pháp range được sử dụng trong vòng lặp for để lặp qua các phần tử của một slices
	func main() {
		for b, c := range a {
			fmt.Printf("2**%d = %d\n", b, c)
		}
	}


	
/////////////////Range continued
	pow := make([]int, 10)
	for i := range pow {
		pow[i] = 1 << uint(i) // << toán tử dịch sang trái
	}
	//index là chỉ số phần tử trong pow[]
	for index, value := range pow {
		fmt.Printf("%d %d\n", index, value)
	}
}
