package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	odd := n - n/2
	fmt.Println(float64(odd) / float64(n))
}
