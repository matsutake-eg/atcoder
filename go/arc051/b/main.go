package main

import "fmt"

func main() {
	var k int
	fmt.Scan(&k)

	a, b := 2, 1
	for i := 1; i < k; i++ {
		b, a = a, a+b
	}
	fmt.Println(a, b)
}
