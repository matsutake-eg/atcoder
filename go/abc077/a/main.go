package main

import "fmt"

func main() {
	var s1, s2 string
	fmt.Scan(&s1, &s2)

	if s1[0] == s2[2] && s1[1] == s2[1] && s1[2] == s2[0] {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
