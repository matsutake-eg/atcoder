package main

import "fmt"

func main() {
	var (
		n int
		s string
	)
	fmt.Scan(&n, &s)

	ans := 1
	for i := 1; i < len(s); i++ {
		if s[i-1] != s[i] {
			ans++
		}
	}
	fmt.Println(ans)
}
