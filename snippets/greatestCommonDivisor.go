package main

import "fmt"

func gcd(x, y int) int {
	for y != 0 {
		x, y = y, x%y
	}
	return x
}

func main() {
	var x, y int
	fmt.Scan(&x, &y)

	fmt.Println(gcd(x, y))
}
