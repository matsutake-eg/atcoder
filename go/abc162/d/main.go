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

	if len(xm) < 3 {
		fmt.Println(0)
		return
	}
	ans := xm['R'] * xm['G'] * xm['B']

	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if s[i] == s[j] {
				continue
			}
			d := j - i
			if k := j + d; k < n && s[i] != s[j] && s[j] != s[k] && s[k] != s[i] {
				ans--
			}
		}
	}
	fmt.Println(ans)
}
