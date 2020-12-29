package main

import "fmt"

func main() {
	var n, m int
	fmt.Scan(&n, &m)

	if m < n {
		fmt.Println("No")
	} else {
		fmt.Println("Yes")
	}
}
