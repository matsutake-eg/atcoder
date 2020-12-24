package main

import "fmt"

func main() {
	var s string
	fmt.Scan(&s)

	if len(s) == 2 {
		fmt.Println(s)
	} else {
		fmt.Println(s[2:3] + s[1:2] + s[0:1])
	}
}
