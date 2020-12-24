package main

import "fmt"

func main() {
	var s string
	fmt.Scan(&s)

	c := rune(s[0])
	ans := 0
	for _, r := range s[1:] {
		if c != r {
			ans++
			c = r
		}
	}
	fmt.Println(ans)
}
