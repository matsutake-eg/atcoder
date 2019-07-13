package main

import "fmt"

func fib(x int) int {
	a, b := 0, 1
	for i := 0; i < x; i++ {
		a, b = b, a+b
	}
	return a
}

func main() {
	var x int
	fmt.Scan(&x)

	fmt.Println(fib(x))
}
