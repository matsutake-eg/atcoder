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
	var x, y int
	fmt.Scan(&x, &y)

	fmt.Println(lcm(x, y))
}
