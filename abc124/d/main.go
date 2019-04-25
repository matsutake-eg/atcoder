package main

import "fmt"

func main() {
	var n, k int
	var s string
	fmt.Scan(&n, &k, &s)

	ans := 0
	for l, r := -1, 0; r < len(s); r++ {
		for r < len(s) && s[r] == '1' {
			r++
		}
		for k > 0 {
			k--
			for r < len(s) && s[r] == '0' {
				if r == len(s) {
					break
				}
				r++
			}
			for r < len(s) && s[r] == '1' {
				if r == len(s) {
					break
				}
				r++
			}
		}
		r--
		if v := r - l; v > ans {
			ans = v
		}
		k++
		l++
		for l < len(s) && s[l] == '1' {
			l++
		}
		for l < len(s) && s[l] == '0' {
			l++
		}
		l--
	}

	fmt.Println(ans)
}
