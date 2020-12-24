package main

import "fmt"

func min(x, y int) int {
	if x < y {
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
	for i := 1; i < n; i++ {
		lm := make(map[rune]bool)
		for _, r := range s[:i] {
			lm[r] = true
		}

		rm := make(map[rune]bool)
		for _, r := range s[i:] {
			if _, ok := lm[r]; ok {
				rm[r] = true
			}
		}
		if v := len(rm); v > ans {
			ans = v
		}
	}

	fmt.Println(ans)
}
