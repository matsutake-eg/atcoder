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
	n := scanInt()
	xs := make([]int, n+1)
	for i := 1; i <= n; i++ {
		x := scanInt()
		m := 0
		switch i {
		case 1:
			m = 100000
		case 2:
			m = 50000
		case 3:
			m = 30000
		case 4:
			m = 20000
		case 5:
			m = 10000
		}
		xs[x] = m
	}

	wr := bufio.NewWriter(os.Stdout)
	for _, v := range xs[1:] {
		wr.WriteString(strconv.Itoa(v) + "\n")
	}
	wr.Flush()
}
