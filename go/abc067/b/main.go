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
	l := make([]int, n)
	for i := range l {
		l[i] = readInt()
	}

	sort.Sort(sort.Reverse(sort.IntSlice(l)))
	ans := 0
	for i := 0; i < k; i++ {
		ans += l[i]
	}
	fmt.Println(ans)
}
