package main

import "fmt"

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func main() {
	var x1, y1, r int
	fmt.Scan(&x1, &y1, &r)
	var x2, y2, x3, y3 int
	fmt.Scan(&x2, &y2, &x3, &y3)

	x1l, x1r := x1-r, x1+r
	y1t, y1b := y1+r, y1-r
	if x1l <= x2 && x3 <= x1r && y1b <= y2 && y3 <= y1t {
		for _, v := range [][]int{{x2, y2}, {x2, y3}, {x3, y2}, {x3, y3}} {
			d2 := (x1-v[0])*(x1-v[0]) + (y1-v[1])*(y1-v[1])
			if d2 > r*r {
				fmt.Println("YES")
				fmt.Println("YES")
				return
			}
		}
		fmt.Println("YES")
		fmt.Println("NO")
	} else if x2 <= x1l && x1r <= x3 && y2 <= y1b && y1t <= y3 {
		fmt.Println("NO")
		fmt.Println("YES")
	} else {
		fmt.Println("YES")
		fmt.Println("YES")
	}
}
