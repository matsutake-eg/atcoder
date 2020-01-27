package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
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

func main() {
	n, k := readInt(), readInt()
	hs := make([]int, n)
	for i := range hs {
		hs[i] = readInt()
	}

	sort.Ints(hs)
	ans := 0
	for i := n - 1; i >= 0; i-- {
		if k <= 0 {
			ans += hs[i]
		} else {
			k--
		}
	}
	fmt.Println(ans)
}
