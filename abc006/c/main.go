package main

import "fmt"

func main() {
	var n, m int
	fmt.Scan(&n, &m)

	for j := 0; j <= n; j++ {
		a := n - j
		b := m - j*3
		if b < 2 || b%2 == 1 {
			continue
		}

		i := (a*4 - b) / 2
		k := (b - a*2) / 2
		if i >= 0 && k >= 0 {
			fmt.Println(i, j, k)
			return
		}
	}

	fmt.Println(-1, -1, -1)
}
