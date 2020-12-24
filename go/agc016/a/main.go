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

func main() {
	var s string
	fmt.Scan(&s)

	ans := math.MaxInt64
	for _, r := range "abcdefghijklmnopqrstuvwxyz" {
		t := s
	L:
		for i := 0; i < len(s); i++ {
			for j := range t {
				if j+1 < len(t) && t[j] == t[j+1] {
					continue
				} else if j < len(t)-1 {
					break
				}
				ans = min(ans, i)
				break L
			}

			xs := make([]rune, len(t)-1)
			for j := range xs {
				if rune(t[j]) == r || rune(t[j+1]) == r {
					xs[j] = r
				} else {
					xs[j] = rune(t[j])
				}
			}
			t = string(xs)
		}
	}
	fmt.Println(ans)
}
