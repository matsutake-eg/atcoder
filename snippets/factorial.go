package main

import "fmt"

func fac(x int) int {
	ans := 1
	for i := x; i > 0; i-- {
		ans *= i
	}
	return ans
}

func main() {
	var x int
	fmt.Scan(&x)

	fmt.Println(fac(x))
}
