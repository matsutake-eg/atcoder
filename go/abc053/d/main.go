package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func scanInt() int {
	sc.Scan()
	iv, _ := strconv.Atoi(sc.Text())
	return iv
}

func init() {
	sc.Split(bufio.ScanWords)
}

func main() {
	n := scanInt()
	xm := make(map[int]bool, n)
	for i := 0; i < n; i++ {
		a := scanInt()
		xm[a] = true
	}

	if v := len(xm); v%2 == 1 {
		fmt.Println(v)
	} else {
		fmt.Println(v - 1)
	}
}
