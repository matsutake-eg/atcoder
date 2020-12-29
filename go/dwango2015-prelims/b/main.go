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
	sc.Buffer(make([]byte, 100001), 100001*100)
}

func main() {
	s := scanString()
	cnt := 0
	ans := 0
	for i := 0; i < len(s)-1; i++ {
		if s[i:i+2] == "25" {
			cnt++
			i++
		} else {
			cnt = 0
		}
		ans += cnt
	}
	fmt.Println(ans)
}
