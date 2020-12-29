package main

import "fmt"

func main() {
	var a, b, k int
	fmt.Scan(&a, &b, &k)

	if k < a {
		a -= k
	} else if k > a+b {
		a, b = 0, 0
	} else {
		k -= a
		a = 0
		b -= k
	}

	fmt.Println(a, b)
}
