package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	if n%2 == 0 {
		fmt.Println(n)
		return
	}
	fmt.Println(n * 2)
}
