package main

import (
	"fmt"
)

func main() {
	var s string
	fmt.Scan(&s)

	ans := ""
	for _, r := range s {
		if r == 'B' {
			if ans != "" {
				ans = ans[:len(ans)-1]
			}
		} else {
			ans += string(r)
		}
	}
	fmt.Println(ans)
}
