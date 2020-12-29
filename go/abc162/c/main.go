package main

import "fmt"

func gcd(x, y int) int {
	for y > 0 {
		x, y = y, x%y
	}
	return x
}

func main() {
	var k int
	fmt.Scan(&k)

	ans := 0
	for a := 1; a <= k; a++ {
		for b := 1; b <= k; b++ {
			for c := 1; c <= k; c++ {
				g := gcd(a, b)
				ans += gcd(g, c)
			}
		}
	}
	fmt.Println(ans)
}
