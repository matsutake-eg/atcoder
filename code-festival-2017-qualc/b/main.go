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
	all := 1
	odd := 1
	for i := 0; i < n; i++ {
		all *= 3
		a := scanInt()
		if a%2 == 0 {
			odd *= 2
		}
	}
	fmt.Println(all - odd)
}
