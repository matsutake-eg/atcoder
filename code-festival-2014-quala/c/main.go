package main

import "fmt"

func main() {
	var a, b int
	fmt.Scan(&a, &b)

	x := b/4 - (a-1)/4
	y := b/100 - (a-1)/100
	z := b/400 - (a-1)/400
	fmt.Println(x - y + z)
}
