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

var pm, qm map[int]bool

func pin(x int) string {
	if pm[x] {
		return "."
	}
	if qm[x] {
		return "o"
	}
	return "x"
}

func main() {
	a, b := scanInt(), scanInt()
	pm = make(map[int]bool)
	for i := 0; i < a; i++ {
		p := scanInt()
		pm[p] = true
	}
	qm = make(map[int]bool)
	for i := 0; i < b; i++ {
		q := scanInt()
		qm[q] = true
	}

	wr := bufio.NewWriter(os.Stdout)
	wr.WriteString(pin(7) + " " + pin(8) + " " + pin(9) + " " + pin(0) + "\n")
	wr.WriteString(" " + pin(4) + " " + pin(5) + " " + pin(6) + "\n")
	wr.WriteString("  " + pin(2) + " " + pin(3) + "\n")
	wr.WriteString("   " + pin(1) + "\n")
	wr.Flush()
}
