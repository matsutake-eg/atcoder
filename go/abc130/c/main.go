package main

import "fmt"

func main() {
	var w, h, x, y float64
	fmt.Scan(&w, &h, &x, &y)

	eq := 0
	if x == w/2 && y == h/2 {
		eq = 1
	}
	fmt.Println(w*h/2, eq)
}
