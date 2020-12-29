package main

import (
	"bufio"
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

func main() {
	h, w, n := scanInt(), scanInt(), scanInt()
	a := make([]int, n)
	for i := range a {
		a[i] = scanInt()
	}

	ans := make([][]int, h)
	for i := range ans {
		ans[i] = make([]int, w)
	}

	x, y, d := -1, 0, 1
	for i, v := range a {
		for j := 0; j < v; j++ {
			if d == 1 {
				x++
				if x == w {
					x--
					y++
					d = -1
				}
			} else {
				x--
				if x == -1 {
					x++
					y++
					d = 1
				}
			}
			ans[y][x] = i + 1
		}
	}

	var wr = bufio.NewWriter(os.Stdout)
	for i := range ans {
		for j := range ans[i] {
			wr.WriteString(strconv.Itoa(ans[i][j]))
			if j < len(ans[i])-1 {
				wr.WriteString(" ")
			} else {
				wr.WriteString("\n")
			}
		}
	}
	wr.Flush()
}
