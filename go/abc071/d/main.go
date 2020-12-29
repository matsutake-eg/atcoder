package main

import "fmt"

func main() {
	var (
		n      int
		s1, s2 string
	)
	fmt.Scan(&n, &s1, &s2)

	const mod = 1000000007
	ans := 0
	i := 0
	if s1[0] == s2[0] {
		ans += 3
		i++
	} else {
		ans += 6
		i += 2
	}
	for i < n {
		if s1[i] == s2[i] {
			if s1[i-1] == s2[i-1] {
				ans *= 2
				ans %= mod
			}
			i++
		} else {
			if s1[i-1] == s2[i-1] {
				ans *= 2
				ans %= mod
			} else {
				ans *= 3
				ans %= mod
			}
			i += 2
		}
	}
	fmt.Println(ans)
}
