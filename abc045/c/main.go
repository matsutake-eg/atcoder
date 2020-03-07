package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var s string
	fmt.Scan(&s)

	ans := 0
	for i := 0; i < 1<<uint(len(s)-1); i++ {
		t := s
		for j := len(s) - 1; j >= 0; j-- {
			b := 1 << uint(j)
			if b == i&b {
				t = t[:j+1] + "+" + t[j+1:]
			}
		}
		for _, v := range strings.Split(t, "+") {
			iv, _ := strconv.Atoi(v)
			ans += iv
		}
	}
	fmt.Println(ans)
}
