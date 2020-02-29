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
	l, h, n := scanInt(), scanInt(), scanInt()
	var wr = bufio.NewWriter(os.Stdout)
	for i := 0; i < n; i++ {
		a := scanInt()
		var x int
		if a > h {
			x = -1
		} else if a <= h && a >= l {
			x = 0
		} else {
			x = l - a
		}
		wr.WriteString(strconv.Itoa(x) + "\n")
	}
	wr.Flush()
}
