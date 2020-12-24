package main

import (
	"fmt"
	"strconv"
)

func main() {
	var s string
	fmt.Scan(&s)
	x, _ := strconv.Atoi(s[0:2])
	y, _ := strconv.Atoi(s[2:4])
	if x >= 1 && x <= 12 && y >= 1 && y <= 12 {
		fmt.Println("AMBIGUOUS")
		return
	}
	if x >= 1 && x <= 12 && (y > 12 || y == 0) {
		fmt.Println("MMYY")
		return
	}
	if (x > 12 || x == 0) && y >= 1 && y <= 12 {
		fmt.Println("YYMM")
		return
	}
	fmt.Println("NA")
}
