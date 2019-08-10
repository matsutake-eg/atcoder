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
	var a, b int
	fmt.Scan(&a, &b)

	fmt.Println(max(a+b, a-b, a*b))
}
