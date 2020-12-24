package main

import "fmt"

func main() {
	var a, b, x int
	fmt.Scan(&a, &b, &x)

	if x < a || x > a+b {
		fmt.Println("NO")
		return
	}
	fmt.Println("YES")
}
