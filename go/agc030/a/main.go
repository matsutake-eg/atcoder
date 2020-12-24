package main

import "fmt"

func main() {
	var a, b, c int
	fmt.Scan(&a, &b, &c)

	if c <= a+b {
		fmt.Println(b + c)
	} else {
		fmt.Println(a + b*2 + 1)
	}
}
