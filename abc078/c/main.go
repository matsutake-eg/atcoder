package main

import "fmt"

func main() {
	var n, m int
	fmt.Scan(&n, &m)

	b := 1
	for i := 0; i < m; i++ {
		b *= 2
	}
	fmt.Println((1900*m + 100*(n-m)) * b)
}
