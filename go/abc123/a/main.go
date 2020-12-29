package main

import "fmt"

func main() {
	var a, b, c, d, e, k int
	fmt.Scan(&a, &b, &c, &d, &e, &k)
	if e-a > k {
		fmt.Println(":(")
		return
	}
	fmt.Println("Yay!")
}
