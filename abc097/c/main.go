package main

import (
	"fmt"
	"sort"
)

func main() {
	var (
		s string
		k int
	)
	fmt.Scan(&s, &k)

	m := make(map[string]bool, len(s)*k)
	xs := make([]string, 0, len(m))
	for i := 1; i <= k; i++ {
		for j := 0; j+i <= len(s); j++ {
			w := s[j : j+i]
			if m[w] {
				continue
			}
			m[w] = true
			xs = append(xs, w)
		}
	}

	sort.Strings(xs)
	fmt.Println(xs[k-1])
}
