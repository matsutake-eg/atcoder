package main

import "fmt"

func main() {
	var a, b, k int
	fmt.Scan(&a, &b, &k)

	for i := 0; i < k; i++ {
		if i%2 == 0 {
			if a%2 == 1 {
				a--
			}
			b += a / 2
			a /= 2
		} else {
			if b%2 == 1 {
				b--
			}
			a += b / 2
			b /= 2
		}
	}
	fmt.Println(a, b)
}
