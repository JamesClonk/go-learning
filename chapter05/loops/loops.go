//#!/usr/bin/env goplay

package main

import (
	"fmt"
)

func main() {
	a := 0
OUTER:
	for {
		b := 0
		for {
			if b == 10 {
				break OUTER
			}
			fmt.Println(a, b)
			b++
		}
		a++
	}
}
