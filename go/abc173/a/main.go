package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	r := n % 1000
	if r == 0 {
		fmt.Println(0)
	} else {
		fmt.Println(1000 - r)
	}
}
