package main

import "fmt"

func main() {
	var w string
	fmt.Scan(&w)

	xm := make(map[rune]int)
	for _, r := range w {
		xm[r]++
	}

	for _, v := range xm {
		if v%2 != 0 {
			fmt.Println("No")
			return
		}
	}
	fmt.Println("Yes")
}
