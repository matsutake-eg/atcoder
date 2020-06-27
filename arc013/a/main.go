package main

import "fmt"

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func main() {
	var n, m, l, p, q, r int
	fmt.Scan(&n, &m, &l, &p, &q, &r)

	ans := 0
	for _, v := range [][]int{{p, q, r}, {p, r, q}, {q, p, r}, {q, r, p}, {r, p, q}, {r, q, p}} {
		x := 1
		for i, bv := range []int{n, m, l} {
			if v[i] > bv {
				x = 0
				break
			}
			x *= bv / v[i]
		}
		ans = max(ans, x)
	}
	fmt.Println(ans)
}
