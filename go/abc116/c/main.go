package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	var h int
	p, cnt := 0, 0
	for i := 0; i < n; i++ {
		fmt.Scan(&h)
		if v := h - p; v >= 0 {
			cnt += v
		}
		p = h
	}
	fmt.Println(cnt)
}
