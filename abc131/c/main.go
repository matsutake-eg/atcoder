package main

import "fmt"

func gcd(a, b int64) int64 {
	r := a % b
	if r == 0 {
		return b
	}
	return gcd(b, r)
}

func main() {
	var a, b, c, d int64
	fmt.Scan(&a, &b, &c, &d)

	dC := b/c - (a-1)/c
	dD := b/d - (a-1)/d
	l := c * d / gcd(c, d)
	dCandD := b/l - (a-1)/l

	fmt.Println(b - a + 1 - (dC + dD - dCandD))
}
