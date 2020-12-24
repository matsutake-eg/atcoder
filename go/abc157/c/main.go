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
	n, m := scanInt(), scanInt()
	xs := make([]int, n)
	for i := range xs {
		xs[i] = -1
	}

	to := 1
	for i := 1; i < n; i++ {
		to *= 10
	}
	from := to*10 - 1
	if n == 1 {
		to = 0
	}

	xm := make(map[int]string)
	for i := 0; i < m; i++ {
		s, c := scanInt(), scanString()
		if v, ok := xm[s]; ok && v != c {
			fmt.Println(-1)
			return
		}
		xm[s] = c
	}

	for i := to; i <= from; i++ {
		str := strconv.Itoa(i)
		found := true
		for s, c := range xm {
			if string(str[s-1]) != c {
				found = false
				continue
			}
		}
		if found {
			fmt.Println(i)
			return
		}
	}
	fmt.Println(-1)
}
