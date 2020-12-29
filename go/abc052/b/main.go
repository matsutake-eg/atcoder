package main

import "fmt"

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func main() {
	var (
		n int
		s string
	)
	fmt.Scan(&n, &s)

	ans := 0
	x := 0
	for _, r := range s {
		if r == 'I' {
			x++
		} else {
			x--
		}
		ans = max(ans, x)
	}
	fmt.Println(ans)
}
