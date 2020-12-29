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
	sum := 0
	xm := make(map[int]int)
	xm[0]++
	for i := 0; i < n; i++ {
		a := scanInt()
		sum += a
		xm[sum]++
	}

	ans := 0
	for _, v := range xm {
		ans += (v * (v - 1)) / 2
	}
	fmt.Println(ans)
}
