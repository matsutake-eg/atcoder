package main

import "fmt"

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func main() {
	var a, d int
	fmt.Scan(&a, &d)

	fmt.Println(max((a+1)*d, a*(d+1)))
}
