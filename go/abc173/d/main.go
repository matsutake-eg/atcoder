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
	a := make([]int, n)
	for i := range a {
		a[i] = scanInt()
	}
	sort.Sort(sort.Reverse(sort.IntSlice(a)))

	ans := a[0]
	cnt := 1
	for i := 1; i < len(a); i++ {
		if cnt >= len(a)-1 {
			break
		}
		ans += a[i]
		cnt++
		if cnt >= len(a)-1 {
			break
		}
		ans += a[i]
		cnt++
	}
	fmt.Println(ans)
}
