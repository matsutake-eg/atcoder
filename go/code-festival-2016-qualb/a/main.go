package main

import "fmt"

func main() {
	var s string
	fmt.Scan(&s)

	c := "CODEFESTIVAL2016"
	ans := 0
	for i := range s {
		if s[i] != c[i] {
			ans++
		}
	}
	fmt.Println(ans)
}
