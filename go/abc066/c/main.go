package main

import (
	"bufio"
	"container/list"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func readInt() int {
	sc.Scan()
	iv, _ := strconv.Atoi(sc.Text())
	return iv
}

func init() {
	sc.Split(bufio.ScanWords)
}

func main() {
	n := readInt()
	qu := list.New()
	for i := 0; i < n; i++ {
		a := readInt()
		if i%2 == 0 {
			qu.PushBack(a)
		} else {
			qu.PushFront(a)
		}
	}

	var wr = bufio.NewWriter(os.Stdout)
	for qu.Len() > 0 {
		var x int
		if n%2 == 0 {
			x = qu.Remove(qu.Front()).(int)
		} else {
			x = qu.Remove(qu.Back()).(int)
		}
		if qu.Len() > 0 {
			wr.WriteString(strconv.Itoa(x) + " ")
		} else {
			wr.WriteString(strconv.Itoa(x) + "\n")
		}
	}
	wr.Flush()
}
