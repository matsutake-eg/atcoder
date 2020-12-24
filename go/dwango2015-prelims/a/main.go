package main

import "fmt"

func main() {
	var n, x int
	fmt.Scan(&n, &x)

	fmt.Println(540*x + 525*(n-x))
}
