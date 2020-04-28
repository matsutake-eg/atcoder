package main

import "fmt"

func main() {
	var s string
	fmt.Scan(&s)

	xm := make(map[int]int)
	sum := 0
	b := 1
	ans := 0
	for i := len(s) - 1; i >= 0; i-- {
		xm[sum]++
		sum += int(s[i]-'0') * b
		sum %= 2019
		ans += xm[sum]
		b = b * 10 % 2019
	}
	fmt.Println(ans)
}
