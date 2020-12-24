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
	n := float64(scanInt())
	r := scanString()
	ans := 0.0
	for _, v := range r {
		switch v {
		case 'A':
			ans += 4
		case 'B':
			ans += 3
		case 'C':
			ans += 2
		case 'D':
			ans++
		}
	}
	fmt.Println(ans / n)
}
