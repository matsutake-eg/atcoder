package main

import "fmt"

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func main() {
	var (
		n, r int
		s    string
	)
	fmt.Scan(&n, &r, &s)

	ans := 0
	for i := n - 1; i >= 0; i-- {
		if s[i] == '.' {
			ans += max(0, i-r+1)
			break
		}
	}
	for i := 0; i < n; i++ {
		if s[i] == '.' {
			ans++
			i += r - 1
		}
	}
	fmt.Println(ans)
}
