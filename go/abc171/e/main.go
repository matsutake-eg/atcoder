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
	s := 0
	a := make([]int, n)
	for i := range a {
		a[i] = scanInt()
		s ^= a[i]
	}

	wr := bufio.NewWriter(os.Stdout)
	for i, v := range a {
		wr.WriteString(strconv.Itoa(s ^ v))
		if i < len(a)-1 {
			wr.WriteString(" ")
		} else {
			wr.WriteString("\n")
		}
	}
	wr.Flush()
}
