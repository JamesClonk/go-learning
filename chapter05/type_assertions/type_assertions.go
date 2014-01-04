//#!/usr/bin/env goplay

package main

import (
	"fmt"
)

func main() {
	var i interface{} = 99
	var s interface{} = []string{"left", "right"}
	printType(i)
	printType(s)

	j := i.(int)
	printType(j)

	if i, ok := i.(int); ok {
		printType(i)
	}

	if s, ok := s.(int); ok {
		printType(s)
	} else {
		fmt.Println("Could not assert [s] of type int")
	}
	if s, ok := s.([]string); ok {
		printType(s)
	}

	// fmt.Println(s.(string)) // <-- this would panic
}

func printType(x interface{}) {
	fmt.Printf("%T -> %v\n", x, x)
}
