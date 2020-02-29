package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func scanInt() int {
	sc.Scan()
	iv, _ := strconv.Atoi(sc.Text())
	return iv
}

func init() {
	sc.Split(bufio.ScanWords)
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func main() {
	w, h, n := scanInt(), scanInt(), scanInt()

	xl, xr := 0, w
	yl, yr := 0, h
	for i := 0; i < n; i++ {
		x, y, a := scanInt(), scanInt(), scanInt()
		switch a {
		case 1:
			xl = max(xl, x)
		case 2:
			xr = min(xr, x)
		case 3:
			yl = max(yl, y)
		case 4:
			yr = min(yr, y)
		}
	}
	if xr > xl && yr > yl {
		fmt.Println((xr - xl) * (yr - yl))
	} else {
		fmt.Println(0)
	}
}
