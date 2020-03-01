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
	n, a, b, k := scanInt(), scanInt(), scanInt(), scanInt()
	xm := make(map[int]bool, n)
	xm[a], xm[b] = true, true
	for i := 0; i < k; i++ {
		p := scanInt()
		if xm[p] {
			fmt.Println("NO")
			return
		}
		xm[p] = true
	}
	fmt.Println("YES")
}
