package main

import "fmt"

func main() {
	var s string
	fmt.Scan(&s)

	ans := 0
	for i, j := 0, len(s)-1; i < j; {
		if s[i] == s[j] {
			i++
			j--
		} else if s[i] != 'x' && s[j] == 'x' {
			ans++
			j--
		} else if s[i] == 'x' && s[j] != 'x' {
			ans++
			i++
		} else {
			fmt.Println(-1)
			return
		}
	}
	fmt.Println(ans)
}
