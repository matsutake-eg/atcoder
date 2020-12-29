package main

import "fmt"

func main() {
	var s string
	fmt.Scan(&s)

	m := make(map[rune]int)
	for _, r := range s {
		m[r]++
		if m[r] >= 3 {
			fmt.Println("No")
			return
		}
	}
	if len(m) == 2 {
		fmt.Println("Yes")
		return
	}
	fmt.Println("No")
}
