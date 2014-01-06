//#!/usr/bin/env goplay

package main

import (
	"fmt"
)

func main() {
	noArgEx()
	argEx()
}

func noArgEx() {
	x, y := 1, 5
	switch {
	case x < y:
		fmt.Println(x)
	case x > y:
		fmt.Println(y)
	}

	// this is what actually happens behind the scenes
	switch true { // <- implicit condition, bool
	case x < y:
		fmt.Println(x)
	case x > y:
		fmt.Println(y)
	}
}

func argEx() {
	switch i := 5; i {
	case 0:
		fmt.Println("0")
	case 5:
		fmt.Println("5")
	default:
		fmt.Println("Default")
	}
}
