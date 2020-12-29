package main

import "fmt"

func gcd(a, b int) int {
	if b > a {
		a, b = b, a
	}
	for {
		r := a % b
		if r == 0 {
			return b
		}
		a = b
		b = r
	}
}

func main() {
	var a, b, k int
	fmt.Scan(&a, &b, &k)

	r := gcd(a, b)
	for i := r; i > 0; i-- {
		if r%i == 0 {
			k--
		}
		if k == 0 {
			fmt.Println(i)
			break
		}
	}
}
