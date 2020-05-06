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
	n := scanInt()
	type music struct {
		s string
		t int
	}
	ms := make([]music, n)
	for i := range ms {
		ms[i].s, ms[i].t = scanString(), scanInt()
	}
	x := scanString()

	isSleep := false
	ans := 0
	for _, m := range ms {
		if m.s == x {
			isSleep = true
			continue
		}

		if isSleep {
			ans += m.t
		}
	}
	fmt.Println(ans)
}
