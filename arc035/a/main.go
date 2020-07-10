package main

import "fmt"

func main() {
	var s string
	fmt.Scan(&s)

	for i := 0; i*2 <= len(s); i++ {
		if r1, r2 := s[i], s[len(s)-1-i]; r1 == '*' || r2 == '*' || r1 == r2 {
			continue
		}
		fmt.Println("NO")
		return
	}
	fmt.Println("YES")
}
