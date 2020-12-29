package main

import "fmt"

func main() {
	var s string
	fmt.Scan(&s)

	fmt.Println(s[:5] + " " + s[6:13] + " " + s[14:])
}
