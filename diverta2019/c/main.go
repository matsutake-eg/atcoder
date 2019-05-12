package main

import "fmt"

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func main() {
	var (
		n int
		s string
	)
	fmt.Scan(&n)
	count, bxa, bx, xa := 0, 0, 0, 0
	for i := 0; i < n; i++ {
		fmt.Scan(&s)
		for j := 0; j < len(s)-1; j++ {
			if s[j] == 'A' && s[j+1] == 'B' {
				count++
			}
		}
		if s[0] == 'B' && s[len(s)-1] == 'A' {
			bxa++
		} else if s[0] == 'B' {
			bx++
		} else if s[len(s)-1] == 'A' {
			xa++
		}
	}

	if bxa == 0 {
		count += min(bx, xa)
	} else if bx == 0 && xa == 0 {
		count += bxa - 1
	} else {
		count += bxa + min(bx, xa)
	}
	fmt.Println(count)
}
