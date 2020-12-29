package main

import "fmt"

func main() {
	var n, m int
	fmt.Scan(&n, &m)

	if n >= 12 {
		n -= 12
	}

	l := 360 * (float64(m) / 60)
	s := 30*float64(n) + 30*(float64(m)/60)
	var ans float64
	if l > s {
		ans = l - s
	} else {
		ans = s - l
	}
	if ans >= 180 {
		ans = 360 - ans
	}
	fmt.Println(ans)
}
