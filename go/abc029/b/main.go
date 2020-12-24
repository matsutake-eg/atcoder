package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var sc = bufio.NewScanner(os.Stdin)

func scanString() string {
	sc.Scan()
	return sc.Text()
}

func init() {
	sc.Split(bufio.ScanWords)
}

func main() {
	ans := 0
	for i := 0; i < 12; i++ {
		s := scanString()
		if strings.ContainsRune(s, 'r') {
			ans++
		}
	}
	fmt.Println(ans)
}
