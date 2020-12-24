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
	h, w := scanInt(), scanInt()
	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			s := scanString()
			if s == "snuke" {
				fmt.Println(string('A'+j) + strconv.Itoa(i+1))
				return
			}
		}
	}
}
