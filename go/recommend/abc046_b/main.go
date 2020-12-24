package main

import "fmt"

func main() {
	var n, k int
	fmt.Scan(&n, &k)

	p := k
	for i := 1; i < n; i++ {
		p *= k - 1
	}

	fmt.Println(p)
}
