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

	w, b := 0, 0
	for i := 0; i < n; i++ {
		if s[i] == '.' {
			w++
		}
	}
	ans := w
	for i := 0; i < n; i++ {
		if s[i] == '#' {
			b++
		} else {
			w--
		}
		ans = min(ans, w+b)
	}
	fmt.Println(ans)
}
