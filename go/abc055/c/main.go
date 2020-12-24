package main

import "fmt"

func main() {
	var n, m int
	fmt.Scan(&n, &m)

	if m < n*2 {
		fmt.Println(m / 2)
	} else {
		fmt.Println(n + (m-n*2)/4)
	}
}
