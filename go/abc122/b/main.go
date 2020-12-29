package main

import "fmt"

func main() {
	var s string
	fmt.Scan(&s)
	var max, cnt int
	for _, v := range s {
		if v == 'A' || v == 'T' || v == 'C' || v == 'G' {
			cnt++
		} else {
			cnt = 0
		}
		if cnt > max {
			max = cnt
		}
	}
	fmt.Println(max)
}
