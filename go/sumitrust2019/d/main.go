package main

import "fmt"

func main() {
	var (
		n int
		s string
	)
	fmt.Scan(&n, &s)

	ans := 0
	for i := 0; i <= 999; i++ {
		x := fmt.Sprintf("%03d", i)
		p := 0
		for xi := range x {
			found := false
			for si := p; si < n; si++ {
				if x[xi] == s[si] {
					found = true
					p = si + 1
					break
				}
			}
			if !found {
				break
			} else if xi == 2 && found {
				ans++
			}
		}
	}
	fmt.Println(ans)
}
