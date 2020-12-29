package main

import "fmt"

func main() {
	var s string
	fmt.Scan(&s)

	n := len(s)
	for i := range s[:n/2] {
		if s[i] != s[n-1-i] {
			fmt.Println("No")
			return
		}
	}
	for i := range s[:(n-1)/2] {
		if s[i] != s[(n-1)/2-1-i] {
			fmt.Println("No")
			return
		}
	}
	for i := (n + 3) / 2; i < n; i++ {
		if s[i] != s[n-1-i] {
			fmt.Println("No")
			return
		}
	}
	fmt.Println("Yes")
}
