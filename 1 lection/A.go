package main

import (
	"fmt"
)

func main() {
	var x1, y1, x2, y2, x, y int
	fmt.Scan(&x1, &y1, &x2, &y2, &x, &y)

	if x >= x1 && x <= x2 { // Плавец выше или ниже
		if y > y2 {
			fmt.Println("N")
		} else if y < y1 {
			fmt.Println("S")
		}
	} else if y >= y1 && y <= y2 { // Плавец слева или справа
		if x < x1 {
			fmt.Println("W")
		} else if x > x2 {
			fmt.Println("E")
		}
	} else { // Плавец в одном из 4 углов
		if x < x1 && y > y2 {
			fmt.Println("NW")
		} else if x > x2 && y > y2 {
			fmt.Println("NE")
		} else if x < x1 && y < y1 {
			fmt.Println("SW")
		} else if x > x2 && y < y1 {
			fmt.Println("SE")
		}
	}
}
