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
	sc.Buffer(make([]byte, 100001), 100001*100)
}

func main() {
	n := scanInt()
	xm := make(map[string]bool)
	xm["TAKAHASHIKUN"] = true
	xm["Takahashikun"] = true
	xm["takahashikun"] = true
	ans := 0
	for i := 0; i < n; i++ {
		w := scanString()
		if i == n-1 {
			w = w[:len(w)-1]
		}
		if xm[w] {
			ans++
		}
	}
	fmt.Println(ans)
}
