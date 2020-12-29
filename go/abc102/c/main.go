package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func main() {
	var n int
	fmt.Scan(&n)

	sc := bufio.NewScanner(os.Stdin)
	sc.Split(bufio.ScanWords)
	b := make([]int, n)
	for i := range b {
		sc.Scan()
		b[i], _ = strconv.Atoi(sc.Text())
		b[i] -= i + 1
	}

	sort.Ints(b)
	m := b[n/2]

	ans := 0
	for _, v := range b {
		ans += abs(v - m)
	}
	fmt.Println(ans)
}
