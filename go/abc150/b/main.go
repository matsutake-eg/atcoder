package main

import "fmt"

func main() {
	var (
		n int
		s string
	)
	fmt.Scan(&n, &s)

	cnt := 0
	for i := 0; i < n-2; i++ {
		if s[i] == 'A' && s[i+1] == 'B' && s[i+2] == 'C' {
			cnt++
		}
	}
	fmt.Println(cnt)
}
