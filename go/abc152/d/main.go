package main

import (
	"fmt"
)

type pair struct{ a, b int }

func convert(x int) pair {
	a, b := x, x%10
	for {
		if a < 10 {
			break
		}
		a /= 10
	}
	return pair{a, b}
}

func main() {
	var n int
	fmt.Scan(&n)

	m := make(map[pair]int)
	for i := 1; i <= n; i++ {
		m[convert(i)]++
	}
	ans := 0
	for i := 1; i <= n; i++ {
		p := convert(i)
		ans += m[pair{p.b, p.a}]
	}
	fmt.Println(ans)
}
