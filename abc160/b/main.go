package main

import "fmt"

func main() {
	var x int
	fmt.Scan(&x)

	a := x / 500
	x -= a * 500
	b := x / 5
	fmt.Println(a*1000 + b*5)
}
