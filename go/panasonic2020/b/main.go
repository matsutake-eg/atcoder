package main

import "fmt"

func main() {
	var h, w int
	fmt.Scan(&h, &w)

	if h == 1 || w == 1 {
		fmt.Println(1)
		return
	}

	a := (h + 1) / 2
	b := h - a
	c := (w + 1) / 2
	d := w - c
	fmt.Println(a*c + b*d)
}
