package main

import "fmt"

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func main() {
	var n, a, b int
	fmt.Scan(&n, &a, &b)

	fmt.Println(max(0, a+b*(n-1)-(b+a*(n-1))+1))
}
