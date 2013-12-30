#!/usr/bin/env goplay

package main

import (
	"fmt"
)

func main() {
	b := []byte("Hello, 世界")

	r := []rune(string(b))
	for i, c := range r {
		fmt.Printf("%d  %-5c %-5s %-5v\n", i, c, string(c), string(c))
	}

	fmt.Println(string(r[7:]))
	fmt.Println(string(r[8:]))
}
