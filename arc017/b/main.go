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
	n, k := scanInt(), scanInt()
	b := scanInt()
	cnt := 1
	ans := 0
	for i := 1; i < n; i++ {
		a := scanInt()
		if a > b {
			cnt++
		} else {
			cnt = 1
		}
		if cnt >= k {
			ans++
		}
		b = a
	}
	if k == 1 {
		fmt.Println(n)
	} else {
		fmt.Println(ans)
	}
}
