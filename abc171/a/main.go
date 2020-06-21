package main

import "fmt"

func main() {
	var a string
	fmt.Scan(&a)

	az := "abcdefghijklmnopqrstuvwxyz"
	for _, r := range az {
		if r == rune(a[0]) {
			fmt.Println("a")
			return
		}
	}
	fmt.Println("A")
}
