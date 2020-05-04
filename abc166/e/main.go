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

	xm := make(map[int]map[int]bool, n)
	ym := make(map[int]map[int]bool, n)
	for i := 1; i <= n; i++ {
		a := scanInt()
		if x := i - a; x > 0 {
			if _, ok := xm[x]; !ok {
				xm[x] = make(map[int]bool)
			}
			xm[x][i] = true
		}
		y := i + a
		if _, ok := ym[y]; !ok {
			ym[y] = make(map[int]bool)
		}
		ym[y][i] = true
	}

	ans := 0
	for x, m := range xm {
		if _, ok := ym[x]; ok {
			ans += len(m) * len(ym[x])
		}
	}
	fmt.Println(ans)
}
