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

	a = abs(a)
	b = abs(b)
	if a == b {
		fmt.Println("Draw")
	} else if a < b {
		fmt.Println("Ant")
	} else {
		fmt.Println("Bug")
	}
}
