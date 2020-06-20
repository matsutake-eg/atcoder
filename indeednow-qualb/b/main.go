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
		for j := 0; i+j < len(t); j++ {
			if s[j] != t[i+j] {
				break
			}
			if i+j == len(t)-1 {
				fmt.Println(len(t) - j - 1)
				return
			}
		}
	}
	fmt.Println(len(t))
}
