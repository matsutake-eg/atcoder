package main

import (
	"fmt"
	"strings"
)

func main() {
	var s string
	fmt.Scan(&s)

	az := "abcdefghijklmnopqrstuvwxyz"
	if len(s) < 26 {
		for _, r := range az {
			if !strings.ContainsRune(s, r) {
				fmt.Println(s + string(r))
				return
			}
		}
	} else {
		for i := len(s) - 1; i >= 1; i-- {
			if s[i-1] < s[i] {
				for _, r := range az {
					if r > rune(s[i-1]) && !strings.ContainsRune(s[:i], r) {
						fmt.Println(s[:i-1] + string(r))
						return
					}
				}
			}
		}
		fmt.Println(-1)
	}
}
