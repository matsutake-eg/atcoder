package main

import "fmt"

func main() {
	var (
		s, t, u string
		a, b    int
	)
	fmt.Scan(&s, &t, &a, &b, &u)

	m := make(map[string]int)
	m[s] = a
	m[t] = b
	if u == s {
		m[s]--
	} else {
		m[t]--
	}
	fmt.Println(m[s], m[t])
}
