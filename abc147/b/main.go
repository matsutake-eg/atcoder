package main

import "fmt"

func main() {
	var s string
	fmt.Scan(&s)

	ans := 0
	for i := range s[:len(s)/2] {
		if s[i] != s[len(s)-1-i] {
			ans++
		}
	}
	fmt.Println(ans)
}
