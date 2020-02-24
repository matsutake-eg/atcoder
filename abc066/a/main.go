package main

import "fmt"

func max(x ...int) int {
	mv := x[0]
	for _, v := range x[1:] {
		if v > mv {
			mv = v
		}
	}
	return mv
}

func main() {
	var a, b, c int
	fmt.Scan(&a, &b, &c)

	fmt.Println(a + b + c - max(a, b, c))
}
