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
	n, m, a, b := scanInt(), scanInt(), scanInt(), scanInt()
	for i := 0; i < m; i++ {
		if n <= a {
			n += b
		}
		c := scanInt()
		n -= c
		if n < 0 {
			fmt.Println(i + 1)
			return
		}
	}
	fmt.Println("complete")
}
