package main

import "fmt"

func main() {
	var k int
	fmt.Scan(&k)

	e := k / 2
	fmt.Println(e * (k - e))
}
