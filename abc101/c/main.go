package main

import "fmt"

func main() {
	var n, k int
	fmt.Scan(&n, &k)

	r := k - 1
	fmt.Println(1 + (n-k+r-1)/r)
}
