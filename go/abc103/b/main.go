package main

import "fmt"

func main() {
	var s, t string
	fmt.Scan(&s, &t)

	for i := 0; i < len(s); i++ {
		if s[len(s)-i:len(s)]+s[0:len(s)-i] == t {
			fmt.Println("Yes")
			return
		}
	}
	fmt.Println("No")
}
