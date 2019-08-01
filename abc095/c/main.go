package main

import "fmt"

func main() {
	var a, b, c, x, y int
	fmt.Scan(&a, &b, &c, &x, &y)

	maxXY := x
	if y > maxXY {
		maxXY = y
	}
	ans := a*x + b*y
	for i := 1; i <= maxXY; i++ {
		if x > 0 {
			x--
		}
		if y > 0 {
			y--
		}
		if v := a*x + b*y + 2*i*c; v < ans {
			ans = v
		}
	}
	fmt.Println(ans)
}
