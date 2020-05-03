package main

import "fmt"

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func main() {
	var x int
	fmt.Scan(&x)

	for i := 0; i <= 100000; i++ {
		a5 := i * i * i * i * i
		b5 := a5 - x
		for j := 0; j*j*j*j*j <= abs(b5); j++ {
			if j*j*j*j*j == abs(b5) {
				if b5 > 0 {
					fmt.Println(i, j)
				} else {
					fmt.Println(i, -j)
				}
				return
			}
		}
	}

	for i := 0; i <= 100000; i++ {
		a5 := -i * i * i * i * i
		b5 := a5 - x
		for j := 0; j*j*j*j*j <= abs(b5); j++ {
			if j*j*j*j*j == abs(b5) {
				if b5 > 0 {
					fmt.Println(-i, j)
				} else {
					fmt.Println(-i, -j)
				}
				return
			}
		}
	}
}
