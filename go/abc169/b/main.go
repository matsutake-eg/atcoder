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
	sort.Ints(a)
	if a[0] == 0 {
		fmt.Println(0)
		return
	}

	ans := 1
	for _, v := range a {
		if v > 1000000000000000000/ans {
			fmt.Println(-1)
			return
		}
		ans *= v
	}
	fmt.Println(ans)
}
