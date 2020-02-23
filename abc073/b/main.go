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
	ans := 0
	for i := 0; i < n; i++ {
		l, r := readInt(), readInt()
		ans += r - l + 1
	}
	fmt.Println(ans)
}
