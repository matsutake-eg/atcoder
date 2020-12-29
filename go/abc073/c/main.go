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

func init() {
	sc.Split(bufio.ScanWords)
}

func main() {
	n := readInt()
	ans := make(map[int]bool, n)
	for i := 0; i < n; i++ {
		a := readInt()
		if _, ok := ans[a]; ok {
			delete(ans, a)
			continue
		}
		ans[a] = true
	}
	fmt.Println(len(ans))
}
