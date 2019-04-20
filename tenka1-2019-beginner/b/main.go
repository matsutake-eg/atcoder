package main

import "fmt"

func main() {
	var n, k int
	var s string
	fmt.Scan(&n, &s, &k)

	c := s[k-1]
	for i := 0; i < n; i++ {
		if s[i] == c {
			fmt.Print(string(c))
		} else {
			fmt.Print("*")
		}
	}
	fmt.Println()
}
