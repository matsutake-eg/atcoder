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
	var a, b, n int
	fmt.Scan(&a, &b, &n)

	l := lcm(a, b)
	fmt.Println(l * ((n + l - 1) / l))
}
