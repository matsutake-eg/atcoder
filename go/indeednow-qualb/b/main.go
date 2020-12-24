package main

import "fmt"

func main() {
	var s, t string
	fmt.Scan(&s, &t)

	if len(s) != len(t) {
		fmt.Println(-1)
		return
	}

	for i := range s {
		p := i
		for j := 0; j < len(t); j++ {
			if s[j] != t[p] {
				break
			}
			if j == len(t)-1 {
				fmt.Println(i)
				return
			}
			p++
			if p == len(t) {
				p = 0
			}
		}
	}
	fmt.Println(-1)
}
