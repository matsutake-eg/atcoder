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
		n, m      int
		name, kit string
	)
	fmt.Scan(&n, &m, &name, &kit)

	nm := make(map[rune]int)
	for _, r := range name {
		nm[r]++
	}
	km := make(map[rune]int)
	for _, r := range kit {
		km[r]++
	}

	ans := 0
	for k, v := range nm {
		if _, ok := km[k]; !ok {
			fmt.Println(-1)
			return
		}

		ans = max(ans, (v+km[k]-1)/km[k])
	}
	fmt.Println(ans)
}
