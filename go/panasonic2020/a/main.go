package main

import "fmt"

func main() {
	var k int
	fmt.Scan(&k)

	xs := []int{1, 1, 1, 2, 1, 2, 1, 5, 2, 2, 1, 5, 1, 2, 1, 14, 1, 5, 1, 5, 2, 2, 1, 15, 2, 2, 5, 4, 1, 4, 1, 51}
	fmt.Println(xs[k-1])
}
