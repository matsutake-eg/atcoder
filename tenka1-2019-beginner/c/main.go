package main

import (
	"bufio"
	"fmt"
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
	sc.Buffer(make([]byte, 100000), 100000000)
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func main() {
	_ = scanInt()
	s := scanString()
	b, w := 0, 0
	for _, r := range s {
		if r == '.' {
			w++
		}
	}
	ans := w
	for _, r := range s {
		if r == '.' {
			w--
		} else {
			b++
		}
		ans = min(ans, b+w)
	}
	fmt.Println(ans)
}
