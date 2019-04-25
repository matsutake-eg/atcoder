package main

import (
	"fmt"
	"strings"
)

func main() {
	var n int
	var s string
	fmt.Scan(&n, &s)

	l := 0
	r := strings.Count(s, ".")
	ans := r + l
	for i := 0; i < n; i++ {
		if s[i] == '#' {
			l++
		} else {
			r--
		}
		if v := r + l; v < ans {
			ans = v
		}
	}
	fmt.Println(ans)
}
