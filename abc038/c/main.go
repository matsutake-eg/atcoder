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
	a := make([]int, n)
	for i := range a {
		a[i] = scanInt()
	}

	b := make([]int, n)
	for i := n - 1; i >= 1; i-- {
		if a[i-1] < a[i] {
			b[i-1] = b[i] + 1
		}
	}

	ans := n
	for _, v := range b {
		ans += v
	}
	fmt.Println(ans)
}
