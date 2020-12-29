package main

import "fmt"

func main() {
	var n, k int
	fmt.Scan(&n, &k)

	if k == 1 {
		fmt.Println(0)
		return
	}
	fmt.Println(n - k)
}
