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
	var a, b, c, d int
	fmt.Scan(&a, &b, &c, &d)

	dC := b/c - (a-1)/c
	dD := b/d - (a-1)/d
	l := lcm(c, d)
	dCandD := b/l - (a-1)/l

	fmt.Println(b - a + 1 - (dC + dD - dCandD))
}
