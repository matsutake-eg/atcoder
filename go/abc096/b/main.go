package main

import "fmt"

func max(x ...int) int {
	ans := x[0]
	for _, v := range x[1:] {
		if v > ans {
			ans = v
		}
	}
	return ans
}

func main() {
	var a, b, c, k int
	fmt.Scan(&a, &b, &c, &k)

	ans := a + b + c
	mx := max(a, b, c)
	ans -= mx
	for i := 0; i < k; i++ {
		mx *= 2
	}
	ans += mx
	fmt.Println(ans)
}
