package main

import "fmt"

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func main() {
	var s string
	fmt.Scan(&s)

	ca, cb, cc := 0, 0, 0
	for _, r := range s {
		switch r {
		case 'a':
			ca++
		case 'b':
			cb++
		case 'c':
			cc++
		}
	}
	if abs(ca-cb) > 1 || abs(cb-cc) > 1 || abs(cc-ca) > 1 {
		fmt.Println("NO")
	} else {
		fmt.Println("YES")
	}
}
