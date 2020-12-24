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
	n, k := scanInt(), scanInt()
	t := make([]int, n)
	for i := range t {
		t[i] = scanInt()
	}
	xs := make([]int, n)
	xs[0], xs[1] = k, k
	xs[2] = t[0] + t[1] + t[2]
	for i := 3; i < n; i++ {
		xs[i] = xs[i-1] - t[i-3] + t[i]
	}
	for i, v := range xs {
		if v < k {
			fmt.Println(i + 1)
			return
		}
	}
	fmt.Println(-1)
}
