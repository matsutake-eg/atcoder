package main

import (
	"bufio"
	"container/list"
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
	buf := make([]byte, 10000)
	sc.Buffer(buf, 10000000)
}

func main() {
	s := scanString()
	ls := list.New()
	for _, r := range s {
		ls.PushBack(string(r))
	}
	rv := false
	q := scanInt()
	for i := 0; i < q; i++ {
		t := scanInt()
		switch t {
		case 1:
			rv = !rv
		case 2:
			f, c := scanInt(), scanString()
			if rv {
				f = 3 - f
			}
			if f == 1 {
				ls.PushFront(c)
			} else {
				ls.PushBack(c)
			}
		}
	}

	var wr = bufio.NewWriter(os.Stdout)
	if rv {
		for e := ls.Back(); e != nil; e = e.Prev() {
			wr.WriteString(e.Value.(string))
		}
	} else {

		for e := ls.Front(); e != nil; e = e.Next() {
			wr.WriteString(e.Value.(string))
		}
	}
	wr.WriteString("\n")
	wr.Flush()
}
