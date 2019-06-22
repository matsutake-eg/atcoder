package main

import "fmt"

func main() {
	var s string
	fmt.Scan(&s)

	for i := range s[:len(s)-1] {
		if s[i] == s[i+1] {
			fmt.Println("Bad")
			return
		}
	}
	fmt.Println("Good")
}
