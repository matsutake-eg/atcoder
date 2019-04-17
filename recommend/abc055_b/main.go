package main

import "fmt"

func main() {
	var n int64
	fmt.Scan(&n)

	p := int64(1)
	for i := int64(1); i <= n; i++ {
		p *= i
		p %= 1000000007
	}
	fmt.Println(p)
}
