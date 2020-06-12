package main

import "fmt"

func main() {
	var (
		n    int
		s, t string
	)
	fmt.Scan(&n, &s, &t)

	if s == t {
		fmt.Println(len(s))
		return
	}

	ans := s + t
	for i := 1; i < n; i++ {
		if s[len(s)-i:] == t[:i] {
			ans = s + t[i:]
		}
	}
	fmt.Println(len(ans))
}
