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
	sc.Buffer(make([]byte, 100001), 100001*100)
}

func main() {
	n := scanInt()
	s := make([]string, n)
	for i := range s {
		s[i] = scanString()
	}

	var wr = bufio.NewWriter(os.Stdout)
	for j := 0; j < n; j++ {
		for i := n - 1; i >= 0; i-- {
			wr.WriteString(string(s[i][j]))
		}
		wr.WriteString("\n")
	}
	wr.Flush()
}
