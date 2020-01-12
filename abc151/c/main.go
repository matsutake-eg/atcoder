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

func readStr() string {
	sc.Scan()
	return sc.Text()
}

func init() {
	sc.Split(bufio.ScanWords)
}

func main() {
	_, m := readInt(), readInt()
	ans := make(map[int]int)
	ac, pr := 0, 0
	for i := 0; i < m; i++ {
		p, s := readInt(), readStr()
		if _, ok := ans[p]; !ok {
			if s == "AC" {
				ans[p] = 0
				ac++
			} else {
				ans[p] = 1
			}
			continue
		}

		if ans[p] > 0 {
			if s == "AC" {
				pr += ans[p]
				ans[p] = 0
				ac++
				continue
			}
			ans[p]++
		}
	}
	fmt.Println(ac, pr)
}
