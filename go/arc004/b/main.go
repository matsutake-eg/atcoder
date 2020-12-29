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
	d := make([]int, n)
	mx := 0
	l := 0
	for i := range d {
		d[i] = scanInt()
		mx += d[i]
		l = max(l, d[i])
	}

	fmt.Println(mx)
	fmt.Println(max(0, l-(mx-l)))
}
