package main

import "fmt"

func main() {
	var x1, y1, x2, y2 int
	fmt.Scan(&x1, &y1, &x2, &y2)

	dx, dy := x2-x1, y2-y1
	fmt.Println(x2-dy, y2+dx, x1-dy, y1+dx)
}
