package main

import "fmt"

func scanStrs(h, w int) (strs [][]rune) {
	strs = make([][]rune, h)
	var s string
	for i := range strs {
		fmt.Scan(&s)
		strs[i] = []rune(s)
	}
	return
}

func countUp(a [][]rune, x, y int) int {
	if x < 0 || x > len(a)-1 || y < 0 || y > len(a[0])-1 {
		return 0
	} else if a[x][y] == '#' {
		return 1
	}
	return 0
}

func main() {
	var h, w int
	fmt.Scan(&h, &w)
	a := scanStrs(h, w)

	var count int
	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			count = 0
			if a[i][j] == '.' {
				count += countUp(a, i-1, j-1)
				count += countUp(a, i-1, j)
				count += countUp(a, i-1, j+1)
				count += countUp(a, i, j-1)
				count += countUp(a, i, j+1)
				count += countUp(a, i+1, j-1)
				count += countUp(a, i+1, j)
				count += countUp(a, i+1, j+1)
				fmt.Print(count)
			} else {
				fmt.Print(string(a[i][j]))
			}
		}
		fmt.Println()
	}
}
