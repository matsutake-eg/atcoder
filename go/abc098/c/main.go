package main

import "fmt"

func main() {
	var (
		n int
		s string
	)
	fmt.Scan(&n, &s)

	e := make([]int, n)
	w := make([]int, n)
	for i, r := range s {
		if i > 0 {
			e[i] = e[i-1]
			w[i] = w[i-1]
		}
		if r == 'E' {
			e[i]++
			continue
		}
		w[i]++
	}

	ans := 1000000
	for i := 0; i < n; i++ {
		lw := 0
		if i > 0 {
			lw = w[i-1]
		}
		le := 0
		if i < n-1 {
			le = e[len(e)-1] - e[i]
		}
		if v := lw + le; v < ans {
			ans = v
		}
	}
	fmt.Println(ans)
}
