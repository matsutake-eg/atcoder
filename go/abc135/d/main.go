package main

import "fmt"

var (
	d   = 1
	ans [][]int
)

func calc(x, i int) {
	for j := range ans[i-1] {
		nx := (x*d + j) % 13
		ans[i][nx] = (ans[i][nx] + ans[i-1][j]) % 1000000007
	}
}

func main() {
	var s string
	fmt.Scan(&s)

	ans = make([][]int, len(s)+1)
	ans[0] = make([]int, 1)
	ans[0][0] = 1
	for i := 1; i <= len(s); i++ {
		if i > 1 {
			d = d * 10 % 13
		}

		ans[i] = make([]int, 13)
		r := s[len(s)-i]
		if r != '?' {
			calc(int(r-'0'), i)
			continue
		}
		for j := 0; j <= 9; j++ {
			calc(j, i)
		}
	}
	fmt.Println(ans[len(s)][5])
}
