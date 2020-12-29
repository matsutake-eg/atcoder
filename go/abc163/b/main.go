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
	n, m := scanInt(), scanInt()
	sum := 0
	for i := 0; i < m; i++ {
		a := scanInt()
		sum += a
	}

	if sum > n {
		fmt.Println(-1)
	} else {
		fmt.Println(n - sum)
	}
}
