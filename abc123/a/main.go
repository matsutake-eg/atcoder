package main

import "fmt"

func main() {
	var a, b, c, d, e, k int
	fmt.Scan(&a, &b, &c, &d, &e, &k)

	if e-a <= k && d-a <= k && c-a <= k && b-a <= k {
		fmt.Println("Yay!")
	} else {
		fmt.Println(":(")
	}
}
