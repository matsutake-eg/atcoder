package main

import "fmt"

func main() {
	xm := make(map[int]int)
	for i := 0; i < 3; i++ {
		var a, b int
		fmt.Scan(&a, &b)
		xm[a]++
		xm[b]++
	}

	if len(xm) != 4 {
		fmt.Println("NO")
	} else {
		c1, c2 := 0, 0
		for _, v := range xm {
			if v == 1 {
				c1++
			} else if v == 2 {
				c2++
			}
		}
		if c1 == 2 && c2 == 2 {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}
	}
}
