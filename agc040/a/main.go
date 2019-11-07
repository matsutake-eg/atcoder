package main

import "fmt"

func sum(x int) int {
	return (1 + x) * x / 2
}

func main() {
	var s string
	fmt.Scan(&s)

	cnt1, cnt2 := 0, 0
	ans := 0
	for i := range s {
		if s[i] == '<' {
			cnt1++
		} else {
			cnt2++
		}

		if i == len(s)-1 || (s[i] == '>' && s[i+1] == '<') {
			max, min := cnt1, cnt2
			if min > max {
				max, min = min, max
			}
			ans += sum(max)
			ans += sum(min - 1)
			cnt1, cnt2 = 0, 0
		}
	}
	fmt.Println(ans)
}
