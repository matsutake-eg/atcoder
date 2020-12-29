package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
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
}

func main() {
	h, w := scanInt(), scanInt()

	var wr = bufio.NewWriter(os.Stdout)
	wr.WriteString(strings.Repeat("#", w+2) + "\n")
	for i := 0; i < h; i++ {
		s := scanString()
		wr.WriteString("#" + s + "#" + "\n")
	}
	wr.WriteString(strings.Repeat("#", w+2) + "\n")
	wr.Flush()
}
