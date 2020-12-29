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
}

func main() {
	n := scanInt()
	xm := make(map[string]int)
	maxCnt, maxName := 0, ""
	for i := 0; i < n; i++ {
		s := scanString()
		xm[s]++
		if xm[s] > maxCnt {
			maxCnt = xm[s]
			maxName = s
		}
	}
	fmt.Println(maxName)
}
