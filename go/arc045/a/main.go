package main

import (
	"bufio"
	"os"
	"strings"
)

var sc = bufio.NewScanner(os.Stdin)

func scanString() string {
	sc.Scan()
	return sc.Text()
}

func init() {
	sc.Split(bufio.ScanLines)
	sc.Buffer(make([]byte, 100000), 100000000)
}

func main() {
	s := scanString()
	xs := strings.Split(s, " ")
	xm := make(map[string]string)
	xm["Left"] = "<"
	xm["Right"] = ">"
	xm["AtCoder"] = "A"
	wr := bufio.NewWriter(os.Stdout)
	for i, v := range xs {
		wr.WriteString(xm[v])
		if i < len(xs)-1 {
			wr.WriteString(" ")
		} else {
			wr.WriteString("\n")
		}
	}
	wr.Flush()
}
