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
	sum, cnt := 0, 0
	for i := 0; i < n; i++ {
		a := scanInt()
		if a > 0 {
			sum += a
			cnt++
		}
	}
	fmt.Println((sum + cnt - 1) / cnt)
}
