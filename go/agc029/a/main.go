package main

import "fmt"

func main() {
	var s string
	fmt.Scan(&s)

	cnt := 0
	ans := 0
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == 'W' {
			cnt++
		} else {
			ans += cnt
		}
	}
	fmt.Println(ans)
}
