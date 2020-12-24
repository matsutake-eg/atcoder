package main

import "fmt"

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func main() {
	var x, y int
	fmt.Scan(&x, &y)

	fmt.Println(max(x, y))
}
