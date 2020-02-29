package main

import "fmt"

func main() {
	var n, k, x, y int
	fmt.Scan(&n, &k, &x, &y)

	if n <= k {
		fmt.Println(x * n)
	} else {
		fmt.Println(x*k + y*(n-k))
	}
}
