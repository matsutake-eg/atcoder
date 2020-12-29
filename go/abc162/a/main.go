package main

import "fmt"

func main() {
	var n string
	fmt.Scan(&n)

	for _, r := range n {
		if r == '7' {
			fmt.Println("Yes")
			return
		}
	}
	fmt.Println("No")
}
