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
	n, s, t, w := scanInt(), scanInt(), scanInt(), scanInt()
	var ans int
	if w >= s && w <= t {
		ans = 1
	} else {
		ans = 0
	}
	for i := 0; i < n-1; i++ {
		a := scanInt()
		w += a
		if w >= s && w <= t {
			ans++
		}
	}
	fmt.Println(ans)
}
