package main

import "fmt"

func main() {
	var (
		s string
		i int
	)
	fmt.Scan(&s, &i)

	fmt.Println(s[i-1 : i])
}
