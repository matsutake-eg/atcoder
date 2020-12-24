package main

import "fmt"

func mins(x ...int) int {
	mv := x[0]
	for _, v := range x[1:] {
		if v < mv {
			mv = v
		}
	}
	return mv
}

func main() {
	var q, h, s, d, n int
	fmt.Scan(&q, &h, &s, &d, &n)

	x := mins(q*4, h*2, s)
	if x*2 < d {
		fmt.Println(x * n)
	} else {
		fmt.Println(d*(n/2) + x*(n%2))
	}
}
