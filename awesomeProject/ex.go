package main

import "golang.org/x/tour/pic"

// go get golang.org/x/tour/pic
func Pic(dx, dy int) [][]uint8 {
	// Khoi tao slices
	pic := make([][]uint8, dy)
	// Duyệt qua từng dòng của hình ảnh
	for y := range pic {
		// Khởi tạo slice để chứa giá trị màu của mỗi pixel trong dòng
		pic[y] = make([]uint8, dx)

		// Tính toán giá trị màu cho từng pixel trong dòng
		for x := range pic[y] {
			// Ví dụ: sử dụng hàm (x+y)/2 để tạo hình ảnh xám
			pic[y][x] = uint8((x + y) / 2)
		}
	}

	return pic
}

func main() {
	pic.Show(Pic)
}
