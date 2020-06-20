package main

import "fmt"

func main() {
	var (
		a string
		b int
	)
	fmt.Scan(&a, &b)
	b--

	fmt.Println(string(a[b%len(a)]))
}
