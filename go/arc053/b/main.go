package main

import (
	"bufio"
	"fmt"
	"os"
)

var sc = bufio.NewScanner(os.Stdin)

func scanString() string {
	sc.Scan()
	return sc.Text()
}

func init() {
	sc.Split(bufio.ScanWords)
	sc.Buffer(make([]byte, 100000), 100000000)
}

func main() {
	s := scanString()
	xm := make(map[rune]int)
	for _, r := range s {
		xm[r]++
	}
	o, e := 0, 0
	for _, v := range xm {
		o += v % 2
		e += v / 2
	}
	if o == 0 {
		fmt.Println(e * 2)
	} else {
		fmt.Println(1 + e/o*2)
	}
}
