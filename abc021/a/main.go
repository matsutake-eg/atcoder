package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	fmt.Println(n)
	for i := 0; i < n; i++ {
		fmt.Println(1)
	}
}
