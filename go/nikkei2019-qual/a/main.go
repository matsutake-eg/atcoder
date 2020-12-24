package main

import "fmt"

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func main() {
	var n, a, b int
	fmt.Scan(&n, &a, &b)

	fmt.Println(min(a, b), max(0, a-(n-b)))
}
