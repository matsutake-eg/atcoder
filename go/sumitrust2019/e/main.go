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
	const mod = 1000000007
	xm := make(map[int]int)
	xm[0] = 3
	ans := 1
	for i := 0; i < n; i++ {
		a := scanInt()
		ans *= xm[a]
		ans %= mod
		xm[a]--
		xm[a+1]++
	}
	fmt.Println(ans)
}
