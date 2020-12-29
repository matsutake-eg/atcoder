package main

import "fmt"

func gcd(x, y int) int {
	for y > 0 {
		x, y = y, x%y
	}
	return x
}

func lcm(x, y int) int {
	return x / gcd(x, y) * y
}

func main() {
	var x int
	fmt.Scan(&x)

	fmt.Println(lcm(x, 360) / x)
}
