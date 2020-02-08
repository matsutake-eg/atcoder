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
	m := make(map[int]int)
	for i := 0; i < n; i++ {
		a := readInt()
		m[a]++
	}

	if k >= len(m) {
		fmt.Println(0)
		return
	}

	b := make([]int, len(m))
	idx := 0
	for _, c := range m {
		b[idx] = c
		idx++
	}

	sort.Ints(b)
	ans := 0
	for i := 0; i < len(m)-k; i++ {
		ans += b[i]
	}
	fmt.Println(ans)
}
