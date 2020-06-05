package main

import "fmt"

func main() {
	var (
		n int
		s string
	)
	fmt.Scan(&n, &s)

	xm := make(map[rune]int)
	for _, r := range s {
		xm[r]++
	}

	const mod = 1000000007
	ans := 1
	for _, v := range xm {
		ans *= v + 1
		ans %= mod
	}
	fmt.Println(ans - 1)
}
