package main

import "fmt"

func main() {
	var n string
	fmt.Scan(&n)

	for _, r := range n {
		if r == '1' {
			fmt.Print("9")
		} else {
			fmt.Print("1")
		}
	}
	fmt.Println()
}
