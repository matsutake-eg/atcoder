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

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func main() {
	n, k := readInt(), readInt()
	ans := 0
	for i := 0; i < n; i++ {
		x := readInt()
		ans += min(x, k-x) * 2
	}
	fmt.Println(ans)
}
