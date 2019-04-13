package main

import "fmt"

func toggle(r rune) rune {
	if r == '0' {
		return '1'
	}
	return '0'
}

func main() {
	var s string
	fmt.Scan(&s)

	sum := 0
	var now_r rune
	for i, r := range s {
		if i == 0 {
			now_r = r
		} else if r == now_r {
			now_r = toggle(r)
			sum++
		} else {
			now_r = r
		}
	}

	fmt.Println(sum)
}
