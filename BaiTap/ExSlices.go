package main

//go get golang.org/x/tour/pic
//go mod init <tên file>
import "golang.org/x/tour/pic"

//tạo ra một slice hai chiều pic có kích thước dyxdx, đại diện cho hình ảnh.
func Pic(dx, dy int) [][]uint8 {
	// Khởi tạo slice để chứa hình ảnh
	pic := make([][]uint8, dy)

	// Duyệt qua từng dòng của hình ảnh
	for y := range pic {
		// Khởi tạo slice để chứa giá trị màu của mỗi pixel trong dòng có độ dài là dx
		pic[y] = make([]uint8, dx)

		// Tính toán giá trị màu cho từng pixel trong dòng
		for x := range pic[y] {
			// sử dụng hàm x*y để tạo hình ảnh
			pic[y][x] = uint8(x * y)
		}
	}

	return pic
}

func main() {
	pic.Show(Pic)
}
