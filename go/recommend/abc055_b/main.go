package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	p := 1
	for i := 1; i <= n; i++ {
		p *= i
		p %= 1000000007
	}
	fmt.Println(p)
}
