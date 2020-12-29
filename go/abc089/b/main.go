package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	var s string
	for i := 0; i < n; i++ {
		fmt.Scan(&s)
		if s == "Y" {
			fmt.Println("Four")
			return
		}
	}
	fmt.Println("Three")
}
