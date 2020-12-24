package main

import "fmt"

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func main() {
	var n, a, b int
	fmt.Scan(&n, &a, &b)

	if (b-a)%2 == 0 {
		fmt.Println((b - a) / 2)
	} else {
		fmt.Println(min(a-1, n-b) + (b-a+1)/2)
	}
}
