package main

import "fmt"

func maxs(x ...int) int {
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

	m := maxs(a, b, c)
	if m%2 == 0 {
		fmt.Println(0)
	} else {
		if a == m {
			fmt.Println(b * c)
		} else if b == m {
			fmt.Println(a * c)
		} else {
			fmt.Println(a * b)
		}
	}
}
