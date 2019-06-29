package main

import "fmt"

var fm map[int]int

func fac(x int) int {
	ans := 1
	for i := x; i > 0; i-- {
		if v, ok := fm[i]; ok {
			fm[x] = ans * v
			return fm[x]
		}
		ans *= i
	}
	fm[x] = ans
	return fm[x]
}

func main() {
	var n int
	fmt.Scan(&n)

	fm = make(map[int]int)
	for i := 1; i <= n; i++ {
		fmt.Println(fac(i))
	}
}
