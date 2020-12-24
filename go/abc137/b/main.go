package main

import "fmt"

func main() {
	var k, x int
	fmt.Scan(&k, &x)

	for i := x - k + 1; i < x+k; i++ {
		if i != x+k-1 {
			fmt.Printf("%d ", i)
			continue
		}
		fmt.Printf("%d\n", i)
	}
}
