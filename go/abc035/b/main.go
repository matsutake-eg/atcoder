package main

import "fmt"

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func main() {
	var (
		s string
		t int
	)
	fmt.Scan(&s, &t)

	x, y, q := 0, 0, 0
	for _, r := range s {
		switch r {
		case 'L':
			x--
		case 'R':
			x++
		case 'U':
			y--
		case 'D':
			y++
		case '?':
			q++
		}
	}

	d := abs(x) + abs(y)
	if t == 1 {
		fmt.Println(d + q)
	} else {
		if d > q {
			fmt.Println(d - q)
		} else {
			fmt.Println((q - d) % 2)
		}
	}
}
