package main

import (
	"bufio"
	"os"
	"sort"
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
	sc.Buffer(make([]byte, 100000), 100000000)
}

func main() {
	n := scanInt()
	wr := bufio.NewWriter(os.Stdout)
	t := []rune("indeednow")
	sort.Slice(t, func(i, j int) bool { return t[i] < t[j] })
	for i := 0; i < n; i++ {
		s := scanString()
		xs := []rune(s)
		sort.Slice(xs, func(i, j int) bool { return xs[i] < xs[j] })
		if string(xs) == string(t) {
			wr.WriteString("YES")
		} else {
			wr.WriteString("NO")
		}
		wr.WriteString("\n")
	}
	wr.Flush()
}
