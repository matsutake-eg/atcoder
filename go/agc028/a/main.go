package main

import "fmt"

func gcd(x, y int) int {
	for y > 0 {
		x, y = y, x%y
	}
	return x
}

func lcm(x, y int) int {
	return x / gcd(x, y) * y
}

func main() {
	var (
		n, m int
		s, t string
	)
	fmt.Scan(&n, &m, &s, &t)

	l := lcm(n, m)
	xm := make(map[int]byte)
	for i := 0; i < n; i++ {
		xm[1+(l/n)*i] = s[i]
	}
	for i := 0; i < m; i++ {
		if v, ok := xm[1+(l/m)*i]; ok && v != t[i] {
			fmt.Println(-1)
			return
		}
	}
	fmt.Println(l)
}
