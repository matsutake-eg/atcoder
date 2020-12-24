package main

import "fmt"

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func main() {
	var n int
	fmt.Scan(&n)
	xs, ys, hs := make([]int, n), make([]int, n), make([]int, n)
	for i := range xs {
		fmt.Scan(&xs[i], &ys[i], &hs[i])
	}

	for i := 0; i <= 100; i++ {
		for j := 0; j <= 100; j++ {
			var H int
			for k := range xs {
				if hs[k] > 0 {
					H = hs[k] + abs(i-xs[k]) + abs(j-ys[k])
					break
				}
			}
			for k := range xs {
				if newH := max(H-abs(i-xs[k])-abs(j-ys[k]), 0); hs[k] != newH {
					break
				} else if k == len(xs)-1 {
					fmt.Println(i, j, H)
					return
				}
			}
		}
	}
}
