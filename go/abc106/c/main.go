package main

import "fmt"

func main() {
	var (
		s string
		k int
	)
	fmt.Scan(&s, &k)

	l := len(s)
	if k < l {
		l = k
	}
	for _, r := range s[:l] {
		if r != '1' {
			fmt.Println(string(r))
			return
		}
	}
	fmt.Println(1)
}
