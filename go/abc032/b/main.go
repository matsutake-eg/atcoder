package main

import "fmt"

func main() {
	var (
		s string
		k int
	)
	fmt.Scan(&s, &k)

	xm := make(map[string]bool)
	for i := 0; i < len(s); i++ {
		if i+k <= len(s) {
			xm[s[i:i+k]] = true
		}
	}
	fmt.Println(len(xm))
}
