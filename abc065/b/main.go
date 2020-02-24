package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func readInt() int {
	sc.Scan()
	iv, _ := strconv.Atoi(sc.Text())
	return iv
}

func init() {
	sc.Split(bufio.ScanWords)
}

func main() {
	n := readInt()
	xm := make(map[int]int, n)
	for i := 1; i <= n; i++ {
		a := readInt()
		xm[i] = a
	}

	b := 1
	ans := 0
	for {
		if _, ok := xm[b]; !ok {
			fmt.Println(-1)
			return
		}
		old := b
		b = xm[b]
		delete(xm, old)
		ans++
		if b == 2 {
			fmt.Println(ans)
			return
		}
	}
}
