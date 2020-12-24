package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	l := 0
	fmt.Println(l)
	var ls string
	fmt.Scan(&ls)
	if ls == "Vacant" {
		return
	}

	r := n - 1
	fmt.Println(r)
	var rs string
	fmt.Scan(&rs)
	if rs == "Vacant" {
		return
	}

	for {
		m := (l + r) / 2
		fmt.Println(m)
		var ms string
		fmt.Scan(&ms)
		if ms == "Vacant" {
			return
		}

		if v := (r - m) % 2; (v == 1 && rs == ms) || (v == 0 && rs != ms) {
			l = m
			ls = ms
		} else if v := (m - l) % 2; (v == 1 && ms == ls) || (v == 0 && ms != ls) {
			r = m
			rs = ms
		}
	}
}
