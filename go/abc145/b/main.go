package main

import "fmt"

func main() {
	var (
		n int
		s string
	)
	fmt.Scan(&n, &s)

	if n%2 != 0 {
		fmt.Println("No")
		return
	}
	for i := range s[:n/2] {
		if s[i] != s[i+n/2] {
			fmt.Println("No")
			return
		}
	}
	fmt.Println("Yes")
}
