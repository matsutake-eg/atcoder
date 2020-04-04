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
	h, _ := scanInt(), scanInt()
	c := make([]string, h)
	for i := range c {
		c[i] = scanString()
	}

	var wr = bufio.NewWriter(os.Stdout)
	for _, v := range c {
		wr.WriteString(v + "\n")
		wr.WriteString(v + "\n")
	}
	wr.Flush()
}
