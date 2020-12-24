package main

import "fmt"

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func main() {
	var a, b, c, d int
	fmt.Scan(&a, &b, &c, &d)

	if b < c || d < a {
		fmt.Println(0)
	} else {
		fmt.Println(min(b, d) - max(a, c))
	}
}
