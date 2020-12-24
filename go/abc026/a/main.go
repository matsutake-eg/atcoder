package main

import "fmt"

func main() {
	var a int
	fmt.Scan(&a)

	x := a / 2
	y := a - x
	fmt.Println(x * y)
}
