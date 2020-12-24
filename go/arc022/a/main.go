package main

import (
	"fmt"
	"strings"
)

func main() {
	var s string
	fmt.Scan(&s)

	i, c := false, false
	for _, r := range strings.ToLower(s) {
		switch r {
		case 'i':
			i = true
		case 'c':
			if i {
				c = true
			}
		case 't':
			if c {
				fmt.Println("YES")
				return
			}
		}
	}
	fmt.Println("NO")
}
