package main

import "fmt"

func main() {
	var a, b int
	fmt.Scan(&a, &b)

	if a < b*2 {
		fmt.Println(0)
		return
	}
	fmt.Println(a - b*2)
}
