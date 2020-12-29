package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func readInt() int {
	sc.Scan()
	iv, _ := strconv.Atoi(sc.Text())
	return iv
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func init() {
	sc.Split(bufio.ScanWords)
}

func main() {
	n := readInt()
	m := make(map[int]map[int]int)
	for i := 1; i <= n; i++ {
		m[i] = make(map[int]int)
		a := readInt()
		for j := 0; j < a; j++ {
			x, y := readInt(), readInt()
			m[i][x] = y
		}
	}

	ans := 0
	for i := 1; i < 1<<uint(n+1); i++ {
		cnt := 0
	L:
		for j := 1; j <= n; j++ {
			if b := 1 << uint(j); i&b == b {
				cnt++
				for x, y := range m[j] {
					if b := 1 << uint(x); i&b/b != y {
						cnt = 0
						break L
					}
				}
			}
		}
		ans = max(ans, cnt)
	}
	fmt.Println(ans)
}
