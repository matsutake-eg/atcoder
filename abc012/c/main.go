package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	ans := 2025 - n
	for i := 1; i <= 9; i++ {
		for j := 1; j <= 9; j++ {
			if i*j == ans {
				fmt.Println(i, "x", j)
			}
		}
	}
}
