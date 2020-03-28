package main

import (
	"bufio"
	"math"
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

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func main() {
	n, m := scanInt(), scanInt()
	type point struct{ x, y int }
	s := make([]point, n)
	for i := range s {
		s[i].x, s[i].y = scanInt(), scanInt()
	}
	c := make([]point, m)
	for i := range c {
		c[i].x, c[i].y = scanInt(), scanInt()
	}

	var wr = bufio.NewWriter(os.Stdout)
	for _, sv := range s {
		mn := math.MaxInt64
		var ans int
		for ci, cv := range c {
			if d := abs(sv.x-cv.x) + abs(sv.y-cv.y); d < mn {
				mn = d
				ans = ci + 1
			}
		}
		wr.WriteString(strconv.Itoa(ans) + "\n")
	}
	wr.Flush()
}
