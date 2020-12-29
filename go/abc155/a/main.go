package main

import "fmt"

func main() {
	var a, b, c int
	fmt.Scan(&a, &b, &c)

	m := make(map[int]bool)
	m[a], m[b], m[c] = true, true, true
	if len(m) == 2 {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
