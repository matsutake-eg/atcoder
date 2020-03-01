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
	n, x := scanInt(), scanInt()
	ans := 0
	for i := 0; i < n; i++ {
		a := scanInt()
		b := 1 << uint(i)
		if x&b == b {
			ans += a
		}
	}
	fmt.Println(ans)
}
