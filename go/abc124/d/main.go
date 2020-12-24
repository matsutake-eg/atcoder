package main

import "fmt"

func main() {
	var (
		n, k int
		s    string
	)
	fmt.Scan(&n, &k, &s)

	l := 0
	sum := 0
	var bls, brs byte = '1', '1'
	ans := 0
	for r := 0; r < n; r++ {
		if brs != s[r] {
			brs = s[r]
			if s[r] == '0' {
				k--
			}
		}
		if k < 0 {
			for l < n {
				if bls != s[l] {
					bls = s[l]
					if s[l] == '1' {
						k++
						break
					}
				}
				l++
				sum--
			}
		}
		sum++
		if sum > ans {
			ans = sum
		}
	}
	fmt.Println(ans)
}
