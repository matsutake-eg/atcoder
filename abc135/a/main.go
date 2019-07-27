package main

import "fmt"

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func main() {
	var a, b int
	fmt.Scan(&a, &b)

	if a > b {
		a, b = b, a
	}
	if v := abs(a - b); v%2 == 0 {
		fmt.Println(a + v/2)
		return
	}
	fmt.Println("IMPOSSIBLE")
}
