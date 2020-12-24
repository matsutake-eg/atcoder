package main

import (
	"bufio"
	"os"
)

var sc = bufio.NewScanner(os.Stdin)

func scanString() string {
	sc.Scan()
	return sc.Text()
}

func init() {
	sc.Split(bufio.ScanWords)
}

func main() {
	c00, c01, c02, c03 := scanString(), scanString(), scanString(), scanString()
	c10, c11, c12, c13 := scanString(), scanString(), scanString(), scanString()
	c20, c21, c22, c23 := scanString(), scanString(), scanString(), scanString()
	c30, c31, c32, c33 := scanString(), scanString(), scanString(), scanString()

	var wr = bufio.NewWriter(os.Stdout)
	wr.WriteString(c33 + " " + c32 + " " + c31 + " " + c30 + "\n")
	wr.WriteString(c23 + " " + c22 + " " + c21 + " " + c20 + "\n")
	wr.WriteString(c13 + " " + c12 + " " + c11 + " " + c10 + "\n")
	wr.WriteString(c03 + " " + c02 + " " + c01 + " " + c00 + "\n")
	wr.Flush()
}
