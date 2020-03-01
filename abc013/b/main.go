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

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func main() {
	var a, b int
	fmt.Scan(&a, &b)

	fmt.Println(mins(abs(a-b), 10-a+b, 10-b+a))
}
