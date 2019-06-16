package main

import "fmt"

func main() {
	var x, a int
	fmt.Scan(&x, &a)

	if x < a {
		fmt.Println(0)
		return
	}
	fmt.Println(10)
}
