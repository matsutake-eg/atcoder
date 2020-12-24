package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func main() {
	var n, k int
	fmt.Scan(&n, &k)
	sc := bufio.NewScanner(os.Stdin)
	sc.Split(bufio.ScanWords)
	h := make([]int, n)
	for i := range h {
		sc.Scan()
		h[i], _ = strconv.Atoi(sc.Text())
	}

	sort.Ints(h)
	min := 1000000000
	for i := 0; i+k-1 < n; i++ {
		if v := h[i+k-1] - h[i]; v < min {
			min = v
		}
	}
	fmt.Println(min)
}
