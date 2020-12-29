package main

import "fmt"

func main() {
	var s int
	fmt.Scan(&s)

	v := 1000000000
	x := (v - s%v) % v
	y := (s + x) / v
	fmt.Println(0, 0, v, 1, x, y)
}
