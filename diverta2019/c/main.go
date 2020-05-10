package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
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

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func main() {
	n := scanInt()
	ans := 0
	cab, ca, cb := 0, 0, 0
	for i := 0; i < n; i++ {
		s := scanString()
		ans += strings.Count(s, "AB")
		if s[0] == 'B' && s[len(s)-1] == 'A' {
			cab++
		} else if s[0] == 'B' {
			cb++
		} else if s[len(s)-1] == 'A' {
			ca++
		}
	}
	if ca > 0 && cab > 0 {
		ca--
		ans++
	}
	if cb > 0 && cab > 0 {
		cb--
		ans++
	}
	if cab > 0 {
		ans += cab - 1
	}
	ans += min(ca, cb)
	fmt.Println(ans)
}
