package main

import "fmt"

func main() {
	var a, b, c, k int
	fmt.Scan(&a, &b, &c, &k)

	if k <= a {
		fmt.Println(k)
	} else if k <= a+b {
		fmt.Println(a)
	} else {
		m := k - a - b
		fmt.Println(a - m)
	}
}
