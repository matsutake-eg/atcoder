package main

import "fmt"

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func main() {
	var n int
	fmt.Scan(&n)
	var (
		ws   = make([]int, n)
		rsum int
	)
	for i := range ws {
		fmt.Scan(&ws[i])
		rsum += ws[i]
	}
	var (
		lsum int
		ans  = 100000
	)
	for _, v := range ws {
		lsum += v
		rsum -= v
		if d := abs(rsum - lsum); d < ans {
			ans = d
		}
	}
	fmt.Println(ans)
}
