package main

import "fmt"

func main() {
	var s string
	fmt.Scan(&s)

	ans := make([]int, 6)
	for _, r := range s {
		switch r {
		case 'A':
			ans[0]++
		case 'B':
			ans[1]++
		case 'C':
			ans[2]++
		case 'D':
			ans[3]++
		case 'E':
			ans[4]++
		case 'F':
			ans[5]++
		}
	}
	fmt.Println(ans[0], ans[1], ans[2], ans[3], ans[4], ans[5])
}
