package main

import "fmt"

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func main() {
	var s string
	fmt.Scan(&s)

	for i := len(s)/2 - 1; i >= 1; i-- {
		found := true
		for j := 0; j < i; j++ {
			if s[j] != s[i+j] {
				found = false
				break
			}
		}
		if found {
			fmt.Println(i * 2)
			return
		}
	}
}
