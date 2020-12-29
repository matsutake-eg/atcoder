package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
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
	sum := make([]int, n)
	ans := 0
	for i := range sum {
		a, b := scanInt(), scanInt()
		sum[i] = a + b
		ans += -b
	}
	sort.Sort(sort.Reverse(sort.IntSlice(sum)))

	for i := 0; i < n; i += 2 {
		ans += sum[i]
	}
	fmt.Println(ans)
}
