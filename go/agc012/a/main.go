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
	a := make([]int, n*3)
	for i := range a {
		a[i] = scanInt()
	}
	sort.Sort(sort.Reverse(sort.IntSlice(a)))

	ans := 0
	for i := 0; i < n; i++ {
		ans += a[i*2+1]
	}
	fmt.Println(ans)
}
