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
	n, k, m, r := scanInt(), scanInt(), scanInt(), scanInt()
	ss := make([]int, n-1)
	for i := range ss {
		ss[i] = scanInt()
	}
	sort.Sort(sort.Reverse(sort.IntSlice(ss)))

	sum := 0
	for i := 0; i < k-1; i++ {
		sum += ss[i]
	}

	if v := r*k - sum; v > m {
		fmt.Println(-1)
	} else if k < n {
		if ss[k-1] >= v {
			fmt.Println(0)
		} else {
			fmt.Println(v)
		}
	} else {
		if v <= 0 {
			fmt.Println(0)
		} else {
			fmt.Println(v)
		}
	}
}
