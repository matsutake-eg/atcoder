package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func scanInt() int {
	iv, _ := strconv.Atoi(scanString())
	return iv
}

func scanString() string {
	sc.Scan()
	return sc.Text()
}

func init() {
	sc.Split(bufio.ScanWords)
	sc.Buffer(make([]byte, 100000), 100000000)
}

func main() {
	n := scanInt()
	xs := make([]string, n)
	ans := 0
	cnt := make([]int, 9)
	for i := range xs {
		xs[i] = scanString()
		for j, r := range xs[i] {
			switch r {
			case 'x':
				ans++
			case 'o':
				if i == 0 || xs[i-1][j] != 'o' {
					cnt[j]++
				}
			}
		}
	}
	for _, v := range cnt {
		ans += v
	}
	fmt.Println(ans)
}
