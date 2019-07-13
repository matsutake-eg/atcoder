package main

import "fmt"

var c [][]int

func initCmb(x int) {
	c = make([][]int, x+1)
	for i := range c {
		c[i] = make([]int, i+1)
		c[i][0], c[i][len(c[i])-1] = 1, 1
		for j := 1; j < len(c[i])-1; j++ {
			c[i][j] = (c[i-1][j-1] + c[i-1][j])
		}
	}
}

func cmb(x, y int) int {
	if y > x {
		return 0
	}
	return c[x][y]
}

func main() {
	var x, y int
	fmt.Scan(&x, &y)

	if x < y {
		x, y = y, x
	}
	initCmb(x)
	fmt.Println(cmb(x, y))
}
