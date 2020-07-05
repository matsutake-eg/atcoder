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
	sc.Buffer(make([]byte, 100000), 100000000)
}

func main() {
	n := scanInt()
	xm := make(map[string]int)
	for i := 0; i < n; i++ {
		s := scanString()
		xm[s]++
	}

	fmt.Println("AC x", xm["AC"])
	fmt.Println("WA x", xm["WA"])
	fmt.Println("TLE x", xm["TLE"])
	fmt.Println("RE x", xm["RE"])
}
