package main

import "fmt"

func main() {
	var a, b, c, k, s, t int
	fmt.Scan(&a, &b, &c, &k, &s, &t)

	if sum := a*s + b*t; s+t < k {
		fmt.Println(sum)
	} else {
		fmt.Println(sum - c*(s+t))
	}
}
