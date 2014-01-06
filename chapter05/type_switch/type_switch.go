//#!/usr/bin/env goplay

package main

import (
	"fmt"
)

type test interface{}

func main() {
	typeSwitch()
}

func typeSwitch() {
	var x test = 1

	switch x.(type) {
	case int:
		fmt.Printf("int: %T  %v \n", x, x)
	case nil:
		fmt.Printf("nil: %T  %v \n", x, x)
	default:
		fmt.Println("Default")
	}
}
