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

	xm := make(map[int]int, n)
	ym := make(map[int]int, n)
	for i := 1; i <= n; i++ {
		a := scanInt()
		if x := i - a; x > 0 {
			xm[x]++
		}
		y := i + a
		ym[y]++
	}

	ans := 0
	for x := range xm {
		ans += xm[x] * ym[x]
	}
	fmt.Println(ans)
}
