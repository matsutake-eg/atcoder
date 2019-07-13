package main

import "fmt"

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func main() {
	var x int
	fmt.Scan(&x)

	fmt.Println(abs(x))
}
