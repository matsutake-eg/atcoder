package main

import (
	"fmt"
	"math"
)

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func main() {
	var a, b, c int
	fmt.Scan(&a, &b, &c)

	if a%2 == 1 || b%2 == 1 || c%2 == 1 {
		fmt.Println(0)
		return
	} else if a == b && b == c && c == a {
		fmt.Println(-1)
		return
	}

	ans := math.MaxInt64
	for _, v := range []int{a + b, b + c, c + a} {
		cnt := 0
		for v%2 == 0 {
			v /= 2
			cnt++
		}
		ans = min(ans, cnt)
	}
	fmt.Println(ans)
}
