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
}

func main() {
	s, n := scanString(), scanInt()
	for i := 0; i < n; i++ {
		l, r := scanInt(), scanInt()
		rs := make([]rune, 0, r-l+1)
		for j := r - 1; j >= l-1; j-- {
			rs = append(rs, rune(s[j]))
		}
		s = s[:l-1] + string(rs) + s[r:]
	}
	fmt.Println(s)
}
