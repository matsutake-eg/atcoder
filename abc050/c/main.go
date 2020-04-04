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
	const mod = 1000000007
	ans := 1
	if n%2 == 0 {
		cnt := 1
		for i := 0; i < n; i += 2 {
			if !(a[i] == cnt && a[i+1] == cnt) {
				fmt.Println(0)
				return
			}

			cnt += 2
			ans *= 2
			ans %= mod
		}
	} else {
		if a[0] != 0 {
			fmt.Println(0)
			return
		}

		cnt := 2
		for i := 1; i < n; i += 2 {
			if !(a[i] == cnt && a[i+1] == cnt) {
				fmt.Println(0)
				return
			}

			cnt += 2
			ans *= 2
			ans %= mod
		}
	}
	fmt.Println(ans)
}
