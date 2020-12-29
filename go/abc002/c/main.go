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

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func main() {
	xa, ya := scanInt(), scanInt()
	xb, yb := scanInt(), scanInt()
	xc, yc := scanInt(), scanInt()

	a, b := xb-xa, yb-ya
	c, d := xc-xa, yc-ya
	fmt.Println(float64(abs(a*d-b*c)) / 2)
}
