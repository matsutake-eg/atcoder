package main

import "fmt"

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func main() {
	var a, b, c, d int
	fmt.Scan(&a, &b, &c, &d)

	if abs(a-c) <= d || (abs(a-b) <= d && abs(b-c) <= d) {
		fmt.Println("Yes")
		return
	}
	fmt.Println("No")
}
