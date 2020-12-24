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

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func main() {
	n, k, m := readInt(), readInt(), readInt()
	sum := 0
	for i := 0; i < n-1; i++ {
		a := readInt()
		sum += a
	}
	ans := max(n*m-sum, 0)
	if ans > k {
		fmt.Println(-1)
		return
	}
	fmt.Println(ans)
}
