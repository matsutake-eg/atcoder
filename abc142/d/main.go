package main

import "fmt"

func main() {
	var a, b int
	fmt.Scan(&a, &b)

	m := make(map[int]bool)
	for i := 2; i*i <= a; i++ {
		for a%i == 0 {
			a /= i
			m[i] = true
		}
	}
	if a > 1 {
		m[a] = true
	}

	ans := 1
	for k := range m {
		if b%k == 0 {
			ans++
		}
	}
	fmt.Println(ans)
}
