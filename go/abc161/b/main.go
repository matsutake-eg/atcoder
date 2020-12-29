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
	n, m := scanInt(), scanInt()
	a := make([]int, n)
	sum := 0
	for i := range a {
		a[i] = scanInt()
		sum += a[i]
	}

	sort.Sort(sort.Reverse(sort.IntSlice(a)))
	for i := 0; i < m; i++ {
		if a[i]*4*m < sum {
			fmt.Println("No")
			return
		}
	}
	fmt.Println("Yes")
}
