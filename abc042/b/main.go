package main

import (
	"bufio"
	"os"
	"sort"
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
	n, _ := scanInt(), scanInt()
	s := make([]string, n)
	for i := range s {
		s[i] = scanString()
	}

	sort.Strings(s)
	var wr = bufio.NewWriter(os.Stdout)
	for _, v := range s {
		wr.WriteString(v)
	}
	wr.WriteString("\n")
	wr.Flush()
}
