package main

import "fmt"

func main() {
	var s string
	fmt.Scan(&s)

	xs := make([]byte, 0, len(s))
	for i := 0; i < len(s)-1; i++ {
		if s[i:i+2] == "BC" {
			xs = append(xs, 'D')
			i++
			continue
		}
		xs = append(xs, s[i])
	}

	cnt := 0
	ans := 0
	for i := 0; i < len(xs); i++ {
		switch xs[i] {
		case 'A':
			cnt++
		case 'B', 'C':
			cnt = 0
		case 'D':
			ans += cnt
		}
	}
	fmt.Println(ans)
}
