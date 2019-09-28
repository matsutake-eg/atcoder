package main

import "fmt"

func main() {
	var x int
	fmt.Scan(&x)

	m := make(map[int]bool)
	m[1] = true
	for i := 2; i*i <= x; i++ {
		for x%i == 0 {
			x /= i
			m[i] = true
		}
	}
	if x > 1 {
		m[x] = true
	}
	fmt.Println(m)
}
