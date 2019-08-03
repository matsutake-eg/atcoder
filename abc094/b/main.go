package main

import "fmt"

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func main() {
	var n, m, x int
	fmt.Scan(&n, &m, &x)

	lc, rc := 0, 0
	for i := 0; i < m; i++ {
		var a int
		fmt.Scan(&a)
		if a < x {
			lc++
		} else {
			rc++
		}
	}
	fmt.Println(min(lc, rc))
}
