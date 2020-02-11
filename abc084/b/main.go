package main

import "fmt"

func main() {
	var (
		a, b int
		s    string
	)
	fmt.Scan(&a, &b, &s)

	for _, r := range s[:a] {
		if r == '-' {
			fmt.Println("No")
			return
		}
	}
	if s[a] != '-' {
		fmt.Println("No")
		return
	}
	for _, r := range s[a+1:] {
		if r == '-' {
			fmt.Println("No")
			return
		}
	}
	fmt.Println("Yes")
}
