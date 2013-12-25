package main

import (
	"fmt"
	"github.com/jamesclonk/go-learning/chapter01/stacker/stack"
)

func main() {
	var haystack stack.Stack
	haystack.Push("hay")
	haystack.Push(-15)
	haystack.Push([]string{"pin", "clip", "needle"})
	haystack.Push(81.52)

	for {
		item, err := haystack.Pop()
		if err != nil {
			break
		}
		fmt.Println(item)
	}

	var haystackf stack.Stack
	haystackf = haystackf.Pushf("hay")
	haystackf = haystackf.Pushf(-15)
	haystackf = haystackf.Pushf([]string{"pin", "clip", "needle"})
	haystackf = haystackf.Pushf(81.52)
	haystackf.Pushf("abc")
	haystackf.Pushf("zzz")

	for {
		var item interface{}
		var err error
		haystackf, item, err = haystackf.Popf()
		if err != nil {
			break
		}
		fmt.Println(item)
	}
}
