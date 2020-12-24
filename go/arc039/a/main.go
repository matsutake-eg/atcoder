package main

import "fmt"

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func main() {
	var a, b int
	fmt.Scan(&a, &b)

	a1 := a % 10
	a2 := a%100 - a1
	a3 := a - a2 - a1

	b1 := b % 10
	b2 := b%100 - b1
	b3 := b - b2 - b1

	ans := a - b
	ans = max(ans, 900+a2+a1-b)
	ans = max(ans, a3+90+a1-b)
	ans = max(ans, a3+a2+9-b)
	ans = max(ans, a-(100+b2+b1))
	ans = max(ans, a-(b3+b1))
	ans = max(ans, a-(b3+b2))
	fmt.Println(ans)
}
