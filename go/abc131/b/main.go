package main

import "fmt"

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func main() {
	var n, l int
	fmt.Scan(&n, &l)

	ans := 0
	d := 1000
	var e int
	for i := 1; i <= n; i++ {
		t := l + i - 1
		ans += t
		if v := abs(t); v < d {
			d = v
			e = t
		}
	}
	ans -= e
	fmt.Println(ans)
}
