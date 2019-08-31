package main

import (
	"fmt"
	"strconv"
)

func main() {
	var a, b int
	fmt.Scan(&a, &b)

	ans := 0
	for i := a; i <= b; i++ {
		s := strconv.Itoa(i)
		if s[0] == s[4] && s[1] == s[3] {
			ans++
		}
	}
	fmt.Println(ans)
}
