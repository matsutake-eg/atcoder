package main

import "fmt"

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func main() {
	var a, b, k, l int
	fmt.Scan(&a, &b, &k, &l)

	x, y := k/l, k%l
	fmt.Println(min(a*y+b*x, b*(x+1)))
}
