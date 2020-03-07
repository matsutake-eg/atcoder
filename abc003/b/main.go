package main

import (
	"fmt"
	"strings"
)

func main() {
	var s, t string
	fmt.Scan(&s, &t)

	ac := "atcoder"
	for i := 0; i < len(s); i++ {
		if s[i] == t[i] {
			continue
		} else if s[i] == '@' && t[i] != '@' && strings.ContainsRune(ac, rune(t[i])) {
			continue
		} else if s[i] != '@' && t[i] == '@' && strings.ContainsRune(ac, rune(s[i])) {
			continue
		}

		fmt.Println("You will lose")
		return
	}
	fmt.Println("You can win")
}
