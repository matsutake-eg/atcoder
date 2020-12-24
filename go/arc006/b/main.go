package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var sc = bufio.NewScanner(os.Stdin)

func scanString() string {
	sc.Scan()
	return sc.Text()
}

func init() {
	sc.Split(bufio.ScanLines)
	sc.Buffer(make([]byte, 100001), 100001*100)
}

func main() {
	s := scanString()
	ss := strings.Split(s, " ")
	l, _ := strconv.Atoi(ss[1])
	xs := make([]string, l+1)
	for i := range xs {
		xs[i] = scanString()
	}

	p := strings.Index(xs[l], "o")
	for i := l - 1; i >= 0; i-- {
		if p-1 >= 0 && xs[i][p-1] == '-' {
			p -= 2
		} else if p+1 < len(xs[i]) && xs[i][p+1] == '-' {
			p += 2
		}
	}
	fmt.Println(p/2 + 1)
}
