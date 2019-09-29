package main

import "fmt"

var a, p []int

func eat(n, x int) int {
	if n == 0 {
		if x <= 0 {
			return 0
		}
		return 1
	} else if x <= 1+a[n-1] {
		return eat(n-1, x-1)
	} else {
		return p[n-1] + 1 + eat(n-1, x-2-a[n-1])
	}
}

func main() {
	var n, x int
	fmt.Scan(&n, &x)

	a = make([]int, n+1)
	a[0] = 1
	for i := range a[:len(a)-1] {
		a[i+1] = a[i]*2 + 3
	}

	p = make([]int, n+1)
	p[0] = 1
	for i := range p[:len(p)-1] {
		p[i+1] = p[i]*2 + 1
	}

	fmt.Println(eat(n, x))
}
