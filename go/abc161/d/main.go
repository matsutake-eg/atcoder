package main

import (
	"container/list"
	"fmt"
)

func main() {
	var k int
	fmt.Scan(&k)

	ls := list.New()
	for i := 1; i <= 9; i++ {
		ls.PushBack(i)
	}

	var x int
	for i := 0; i < k; i++ {
		x = ls.Remove(ls.Front()).(int)
		d := x % 10
		if d > 0 {
			ls.PushBack(10*x + d - 1)
		}
		ls.PushBack(10*x + d)
		if d < 9 {
			ls.PushBack(10*x + d + 1)
		}
	}
	fmt.Println(x)
}
