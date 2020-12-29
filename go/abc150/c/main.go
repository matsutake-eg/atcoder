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
	p := make([]int, n)
	for i := range p {
		fmt.Scan(&p[i])
	}
	q := make([]int, n)
	for i := range q {
		fmt.Scan(&q[i])
	}

	r := make([]int, n-1)
	fac := 1
	for i := 1; i < n; i++ {
		fac *= i
		r[n-i-1] = fac
	}

	a := 1
	b := 1
	for i := 0; i < n-1; i++ {
		for j := i + 1; j < n; j++ {
			if p[i] > p[j] {
				a += r[i]
			}
			if q[i] > q[j] {
				b += r[i]
			}
		}
	}
	fmt.Println(abs(a - b))
}
