package main

import "fmt"

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func main() {
	var x, t int
	fmt.Scan(&x, &t)

	fmt.Println(max(0, x-t))
}
