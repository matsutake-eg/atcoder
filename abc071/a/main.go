package main

import "fmt"

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func main() {
	var x, a, b int
	fmt.Scan(&x, &a, &b)

	if abs(x-a) < abs(x-b) {
		fmt.Println("A")
	} else {
		fmt.Println("B")
	}
}
