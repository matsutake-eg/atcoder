package main

import "fmt"

func main() {
	var s string
	fmt.Scan(&s)

	ans := 0
	for _, r := range s {
		if r == '+' {
			ans++
		} else {
			ans--
		}
	}
	fmt.Println(ans)
}
