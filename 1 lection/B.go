package main

import (
	"fmt"
)

func main() {
	var x1, y1, x2, y2 int
	var a [2]int
	fmt.Scan(&x1, &y1, &x2, &y2)
	if x1 == 0 && y1 == 0 && y2 == 0 && y2 == 0 {
		fmt.Println(0, 0)
	}
	if x1 == 0 {
		fmt.Println(1, x2+1)
	} else if y1 == 0 {
		fmt.Println(1, y2+1)
	} else {
		if x2 == 0 {
			fmt.Println(x1+1, 1)
		} else if y2 == 0 {
			fmt.Println(y1+1, 1)
		} else {
			sinm := y1 + 1
			krasm := x1 + 1

			sinnosk := y2 + 1
			krasnosk := x2 + 1

			l1 := krasnosk + krasm
			l2 := sinm + sinnosk

			if l1 > l2 {
				a[0] = sinm
				a[1] = sinnosk
			} else {
				a[0] = krasm
				a[1] = krasnosk
			}

			if sinm > max(x1, y1) {
				if (a[0] + a[1]) > sinm+1 {
					a[0] = sinm
					a[1] = 1
				}
			}
			if krasm > max(x1, y1) {
				if (a[0] + a[1]) > krasm+1 {
					a[0] = krasm
					a[1] = 1
				}
			}
			if sinnosk > max(x2, y2) {
				if (a[0] + a[1]) > sinnosk+1 {
					a[0] = 1
					a[1] = sinnosk
				}
			}
			if krasnosk > max(x2, y2) {
				if (a[0] + a[1]) > krasnosk+1 {
					a[0] = 1
					a[1] = krasnosk
				}
			}

			fmt.Println(a[0], a[1])

		}
	}
}
