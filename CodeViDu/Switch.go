package main

import (
	"fmt"
	"time"
)

//	func main() {
//		fmt.Print("Go runs on ")
//		switch os := runtime.GOOS; os {
//		case "darwin":
//			fmt.Println("OS X.")
//		case "linux":
//			fmt.Println("Linux.")
//		default:
//			// freebsd, openbsd,
//			// plan9, windows...
//			fmt.Printf("s%.\n", os)
//		}
//	}

/////////////////////////////////

// func sum(x int64) int64 {

// 	return x + 1
// }

// func main() {
// 	a := sum(2)

// 	switch a {
// 	case 3:
// 		fmt.Print("true")
// 	default:
// 		fmt.Print("false")

// 	}

// }
////////////////////////Switch evaluation order
// func main() {
// 	fmt.Println("When is Wednesday ?")
// 	today := time.Now().Weekday()
// 	switch time.Wednesday {
// 	case today + 0:
// 		fmt.Print("To day")
// 	case today + 1:
// 		fmt.Print("In tomorrow")

// 	case today + 2:
// 		fmt.Print("too far aways")
// 	}

// }

/////////////////Switch with no condition

func main() {
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Sprint("Sáng")
	case t.Hour() < 17:
		fmt.Print("Trưa")
	default:
		fmt.Print("Tối")
	}

}
