package main

import "fmt"

func main() {
	var n string
	fmt.Scan(&n)

	if n[0] == n[2] {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
