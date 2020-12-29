package main

import "fmt"

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func main() {
	var a, b int
	fmt.Scan(&a, &b)

	fmt.Println(max(a, b))
}
