package main

import "fmt"

func main() {
	xm := make(map[int]bool)
	for i := 0; i < 4; i++ {
		var n int
		fmt.Scan(&n)
		xm[n] = true
	}

	if xm[1] && xm[9] && xm[7] && xm[4] {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
