package main

import (
	"fmt"
	"strings"
)

func main() {
	var s string
	fmt.Scan(&s)

	ss := strings.Split(s, "+")
	ans := 0
	for _, v := range ss {
		if !strings.ContainsRune(v, '0') {
			ans++
		}
	}
	fmt.Println(ans)
}
