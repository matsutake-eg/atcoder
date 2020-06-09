package main

import "fmt"

func main() {
	var s string
	fmt.Scan(&s)

	ans := 0
	for i := range s {
		u := len(s) - i - 1
		d := i
		if s[i] == 'U' {
			ans += u + d*2
		} else {
			ans += u*2 + d
		}
	}
	fmt.Println(ans)
}
