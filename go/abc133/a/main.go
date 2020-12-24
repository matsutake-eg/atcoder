package main

import "fmt"

func main() {
	var n, a, b int
	fmt.Scan(&n, &a, &b)

	if v := n * a; v < b {
		fmt.Println(v)
		return
	}
	fmt.Println(b)
}
