package main

import "fmt"

func main() {
	var n, m int
	fmt.Scan(&n, &m)

	if n == 1 && m == 1 {
		fmt.Println(1)
		return
	}
	if n == 1 && m > 1 {
		fmt.Println(m - 2)
		return
	}
	if n > 1 && m == 1 {
		fmt.Println(n - 2)
		return
	}
	fmt.Println((n - 2) * (m - 2))
}
