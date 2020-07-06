package main

import (
	"bufio"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func scanInt() int {
	iv, _ := strconv.Atoi(scanString())
	return iv
}

func scanString() string {
	sc.Scan()
	return sc.Text()
}

func init() {
	sc.Split(bufio.ScanWords)
	sc.Buffer(make([]byte, 100000), 100000000)
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func main() {
	n, m := scanInt(), scanInt()
	b := make([][]int, n)
	ans := make([][]int, n)
	for i := range b {
		b[i] = make([]int, m)
		ans[i] = make([]int, m)
		s := scanString()
		for j, r := range s {
			ri, _ := strconv.Atoi(string(r))
			b[i][j] = ri
		}
	}

	for i := 1; i < len(b)-1; i++ {
		for j := 1; j < len(b[0])-1; j++ {
			mn := b[i-1][j]
			mn = min(mn, b[i+1][j])
			mn = min(mn, b[i][j-1])
			mn = min(mn, b[i][j+1])
			if mn > 0 {
				ans[i][j] = mn
				b[i-1][j] -= mn
				b[i+1][j] -= mn
				b[i][j-1] -= mn
				b[i][j+1] -= mn
			}
		}
	}

	wr := bufio.NewWriter(os.Stdout)
	for _, row := range ans {
		for _, v := range row {
			wr.WriteString(strconv.Itoa(v))
		}
		wr.WriteString("\n")
	}
	wr.Flush()
}
