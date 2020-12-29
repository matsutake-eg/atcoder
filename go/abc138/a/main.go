package main

import "fmt"

func main() {
	var (
		a int
		s string
	)
	fmt.Scan(&a, &s)

	if a >= 3200 {
		fmt.Println(s)
		return
	}
	fmt.Println("red")
}
