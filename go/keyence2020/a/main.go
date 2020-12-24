package main

import "fmt"

func main() {
	var h, w, n int
	fmt.Scan(&h, &w, &n)

	if h < w {
		h, w = w, h
	}
	fmt.Println((n + h - 1) / h)
}
