package main

import "fmt"

func main() {
	var n, k float64
	fmt.Scan(&n, &k)

	fmt.Println(((k-1)*(n-k)*6 + (n-1)*3 + 1) / (n * n * n))
}
