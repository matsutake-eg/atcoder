package main

import "fmt"

func gcd(x, y int) int {
	for y != 0 {
		x, y = y, x%y
	}
	return x
}

func main() {
	var a, b int
	fmt.Scan(&a, &b)

	g := gcd(a, b)
	ans := make(map[int]bool)
	for i := 2; i*i <= g; i++ {
		for g%i == 0 {
			g /= i
			ans[i] = true
		}
	}
	if g > 1 {
		ans[g] = true
	}
	fmt.Println(len(ans) + 1)
}
