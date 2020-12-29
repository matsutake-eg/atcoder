package main

import "fmt"

func lcm(x, y int) int {
	gcd := func(x, y int) int {
		for y != 0 {
			x, y = y, x%y
		}
		return x
	}
	return x * y / gcd(x, y)
}

func main() {
	var a, b int
	fmt.Scan(&a, &b)

	fmt.Println(lcm(a, b))
}
