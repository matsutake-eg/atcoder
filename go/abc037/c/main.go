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
	n, k := scanInt(), scanInt()
	sum := make([]int, n)
	sum[0] = scanInt()
	for i := 1; i < n; i++ {
		x := scanInt()
		sum[i] = sum[i-1] + x
	}

	ans := sum[k-1]
	for i := k; i < n; i++ {
		ans += sum[i] - sum[i-k]
	}
	fmt.Println(ans)
}
