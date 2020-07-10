package main

import "fmt"

func main() {
	var (
		s          string
		a, b, c, d int
	)
	fmt.Scan(&s, &a, &b, &c, &d)

	fmt.Println(s[:a] + "\"" + s[a:b] + "\"" + s[b:c] + "\"" + s[c:d] + "\"" + s[d:])
}
