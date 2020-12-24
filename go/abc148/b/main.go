package main

import "fmt"

func main() {
	var (
		n    int
		s, t string
	)
	fmt.Scan(&n, &s, &t)

	for i := 0; i < n; i++ {
		fmt.Print(string(s[i]))
		fmt.Print(string(t[i]))
	}
	fmt.Println()
}
