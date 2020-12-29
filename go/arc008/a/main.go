package main

import "fmt"

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func main() {
	var n int
	fmt.Scan(&n)

	fmt.Println(min(100*(n/10)+15*(n%10), 100*((n+9)/10)))
}
