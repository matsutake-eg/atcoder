package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func scanInt() int {
	sc.Scan()
	iv, _ := strconv.Atoi(sc.Text())
	return iv
}

func init() {
	sc.Split(bufio.ScanWords)
}

func main() {
	s1, e1 := scanInt(), scanInt()
	s2, e2 := scanInt(), scanInt()
	s3, e3 := scanInt(), scanInt()

	fmt.Println(s1*e1/10 + s2*e2/10 + s3*e3/10)
}
