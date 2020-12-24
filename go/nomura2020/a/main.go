package main

import "fmt"

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func main() {
	var h1, m1, h2, m2, k int
	fmt.Scan(&h1, &m1, &h2, &m2, &k)

	x := h2*60 + m2 - (h1*60 + m1)
	fmt.Println(max(x-k, 0))
}
