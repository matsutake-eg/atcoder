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
	n, h, w := scanInt(), scanInt(), scanInt()
	ans := 0
	for i := 0; i < n; i++ {
		a, b := scanInt(), scanInt()
		if a >= h && b >= w {
			ans++
		}
	}
	fmt.Println(ans)
}
