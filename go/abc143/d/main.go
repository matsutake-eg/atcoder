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
	n := readInt()
	l := make([]int, n)
	for i := range l {
		l[i] = readInt()
	}

	sort.Ints(l)
	ans := 0
	for i := range l {
		for j := i + 1; j < len(l); j++ {
			if idx := sort.SearchInts(l, l[i]+l[j]) - 1; idx > j {
				ans += idx - j
			}
		}
	}
	fmt.Println(ans)
}
