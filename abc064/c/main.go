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

func main() {
	n := scanInt()
	xm := make(map[int]bool)
	cnt := 0
	for i := 0; i < n; i++ {
		a := scanInt()
		if a < 400 {
			xm[1] = true
		} else if a < 800 {
			xm[2] = true
		} else if a < 1200 {
			xm[3] = true
		} else if a < 1600 {
			xm[4] = true
		} else if a < 2000 {
			xm[5] = true
		} else if a < 2400 {
			xm[6] = true
		} else if a < 2800 {
			xm[7] = true
		} else if a < 3200 {
			xm[8] = true
		} else {
			cnt++
		}
	}
	fmt.Println(max(len(xm), 1), len(xm)+cnt)
}
