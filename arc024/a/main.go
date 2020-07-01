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

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func main() {
	L, R := scanInt(), scanInt()
	lm := make(map[int]int)
	for i := 0; i < L; i++ {
		l := scanInt()
		lm[l]++
	}
	rm := make(map[int]int)
	for i := 0; i < R; i++ {
		r := scanInt()
		rm[r]++
	}

	ans := 0
	for k, v := range lm {
		ans += min(v, rm[k])
	}
	fmt.Println(ans)
}
