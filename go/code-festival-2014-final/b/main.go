package main

import (
	"fmt"
	"strconv"
)

func main() {
	var s string
	fmt.Scan(&s)

	ans := 0
	for i, r := range s {
		rv, _ := strconv.Atoi(string(r))
		if i%2 == 0 {
			ans += rv
		} else {
			ans -= rv
		}
	}
	fmt.Println(ans)
}
