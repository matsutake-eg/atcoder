package main

import "fmt"

func main() {
	var a, b, c, k int
	fmt.Scan(&a, &b, &c, &k)

	if k%2 == 1 {
		fmt.Println(b - a)
	} else {
		fmt.Println(a - b)
	}
}
