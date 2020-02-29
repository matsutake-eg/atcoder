package main

import "fmt"

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func main() {
	var w, a, b int
	fmt.Scan(&w, &a, &b)

	if a+w < b || a > b+w {
		fmt.Println(min(abs(b-a-w), abs(a-b-w)))
	} else {
		fmt.Println(0)
	}
}
