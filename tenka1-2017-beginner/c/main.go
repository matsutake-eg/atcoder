package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	for i := 1; i <= 3500; i++ {
		for j := 1; j <= 3500; j++ {
			x := n * i * j
			y := 4*i*j - n*i - n*j
			if y > 0 && x%y == 0 {
				fmt.Println(i, j, x/y)
				return
			}
		}
	}
}
