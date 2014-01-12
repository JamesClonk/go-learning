//#!/usr/bin/env goplay

package main

import (
	"fmt"
)

func main() {
	something("James")("Monday")
	something("Jim")("Tuesday")
}

func something(name string) func(string) {
	return func(day string) {
		fmt.Printf("Hello %s, today is %s\n", name, day)
	}
}
