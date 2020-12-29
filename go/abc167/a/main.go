package main

import "fmt"

func main() {
	var s, t string
	fmt.Scan(&s, &t)

	for i := range s {
		if s[i] != t[i] {
			fmt.Println("No")
			return
		}
	}
	fmt.Println("Yes")
}
