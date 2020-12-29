package main

import (
	"fmt"
	"strings"
)

func main() {
	var s, t string
	fmt.Scan(&s, &t)

	for i := len(s); i >= len(t); i-- {
		found := true
		for j := range t {
			if v := s[i-len(t)+j]; v != t[j] && v != '?' {
				found = false
				break
			}
		}
		if found {
			fmt.Println(strings.Replace(s[:i-len(t)]+t+s[i:], "?", "a", -1))
			return
		}
	}
	fmt.Println("UNRESTORABLE")
}
