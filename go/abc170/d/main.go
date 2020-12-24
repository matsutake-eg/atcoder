package main

import (
	"bufio"
	"fmt"
	"os"
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
	xs := make([]int, 1000001)
	for i := 0; i < n; i++ {
		a[i] = scanInt()
		if xs[a[i]] == 0 {
			for j := a[i]; j < len(xs); j += a[i] {
				xs[j]++
			}
			continue
		}
		xs[a[i]]++
	}

	ans := 0
	for _, v := range a {
		if xs[v] == 1 {
			ans++
		}
	}
	fmt.Println(ans)
}
