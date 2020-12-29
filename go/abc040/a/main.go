package main

import "fmt"

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func main() {
	var n, x int
	fmt.Scan(&n, &x)

	fmt.Println(min(n-x, x-1))
}
