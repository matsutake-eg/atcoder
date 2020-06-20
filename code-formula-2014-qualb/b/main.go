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
	n := scanString()
	o, e := 0, 0
	cnt := 1
	for i := len(n) - 1; i >= 0; i-- {
		x := int(n[i] - '0')
		if cnt%2 == 1 {
			o += x
		} else {
			e += x
		}
		cnt++
	}
	fmt.Println(e, o)
}
