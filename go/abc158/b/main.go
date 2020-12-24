package main

import "fmt"

func main() {
	var n, a, b int
	fmt.Scan(&n, &a, &b)

	d, r := n/(a+b), n%(a+b)
	if r <= a {
		fmt.Println(d*a + r)
	} else {
		fmt.Println(d*a + a)
	}
}
