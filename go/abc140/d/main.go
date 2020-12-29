package main

import "fmt"

func main() {
	var (
		n, k int
		s    string
	)
	fmt.Scan(&n, &k, &s)

	happy := 0
	for i := 1; i < n; i++ {
		if s[i-1] == s[i] {
			happy++
		}
	}
	if v := happy + k*2; v < n {
		fmt.Println(v)
		return
	}
	fmt.Println(n - 1)
}
