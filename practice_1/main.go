package main

import "fmt"

func main() {
	var a int
	fmt.Scan(&a)
	var b, c int
	fmt.Scan(&b, &c)
	var s string
	fmt.Scan(&s)

	fmt.Println(a+b+c, s)
}
