package main

import "fmt"

func main() {
	var s string
	fmt.Scan(&s)

	if s == "keyence" {
		fmt.Println("YES")
		return
	}

	for i := 0; i < len(s); i++ {
		for j := i + 1; j < len(s); j++ {
			if s[:i]+s[j:] == "keyence" {
				fmt.Println("YES")
				return
			}
		}
	}
	fmt.Println("NO")
}
