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
	ans := 0
	xm := make(map[int]bool)
	for i := 0; i < n; i++ {
		a := scanInt()
		if xm[a] {
			ans++
		} else {
			xm[a] = true
		}
	}
	fmt.Println(ans)
}
