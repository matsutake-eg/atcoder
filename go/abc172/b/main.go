package main

import (
	"bufio"
	"fmt"
	"os"
)

var sc = bufio.NewScanner(os.Stdin)

func scanString() string {
	sc.Scan()
	return sc.Text()
}

func init() {
	sc.Split(bufio.ScanWords)
	sc.Buffer(make([]byte, 100000), 100000000)
}

func main() {
	s, t := scanString(), scanString()
	ans := 0
	for i := range s {
		if s[i] != t[i] {
			ans++
		}
	}
	fmt.Println(ans)
}
