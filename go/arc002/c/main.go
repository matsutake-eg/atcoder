package main

import (
	"fmt"
	"math"
)

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

var (
	n   int
	c   string
	ans int = math.MaxInt64
)

func dfs(s string) {
	if len(s) == 4 {
		l, r := s[:2], s[2:]
		cnt := 0
		for i := 0; i < n; i++ {
			cnt++
			if i < n-1 {
				if c[i:i+2] == l || c[i:i+2] == r {
					i++
				}
			}
		}
		ans = min(ans, cnt)
		return
	}
	for _, r := range "ABXY" {
		dfs(s + string(r))
	}
}

func main() {
	fmt.Scan(&n, &c)

	dfs("")
	fmt.Println(ans)
}
