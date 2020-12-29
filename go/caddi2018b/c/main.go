package main

import "fmt"

func prmFac(x int) map[int]int {
	pfm := make(map[int]int)
	for i := 2; i*i <= x; i++ {
		for x%i == 0 {
			x /= i
			pfm[i]++
		}
	}
	if x > 1 {
		pfm[x]++
	}
	return pfm
}

func main() {
	var n, p int
	fmt.Scan(&n, &p)

	ans := 1
	for pn, cnt := range prmFac(p) {
		for i := 0; i < cnt; i++ {
			x := 1
			for j := 0; j < n; j++ {
				x *= pn
				if x > p {
					break
				} else if j == n-1 && p%x == 0 {
					ans *= pn
					p /= x
				}
			}
		}
	}
	fmt.Println(ans)
}
