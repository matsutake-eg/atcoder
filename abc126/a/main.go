package main

import "fmt"

func main() {
	var (
		n, k int
		s    string
	)
	fmt.Scan(&n, &k, &s)
	for i := 0; i < k-1; i++ {
		fmt.Printf("%c", s[i])
	}
	switch s[k-1] {
	case 'A':
		fmt.Print("a")
	case 'B':
		fmt.Print("b")
	case 'C':
		fmt.Print("c")
	}
	for i := k; i < n; i++ {
		fmt.Printf("%c", s[i])
	}
	fmt.Println()
}
