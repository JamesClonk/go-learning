//#!/usr/bin/env goplay

package main

import (
	"fmt"
	"log"
)

func main() {
	something("Let's begin..")
	something("Hello", ",", " ", "World!")
	something("Hello", []string{",", " ", "World!"}...)

	yayForOptions()
	yayForOptions(Option{LogFlag: 3})
	yayForOptions(Option{LogPrefix: "--- ", LogFlag: 3})
}

func something(a string, b ...string) {
	fmt.Print(a)

	for _, v := range b {
		fmt.Print(v)
	}

	fmt.Println()
}

type Option struct {
	LogPrefix string
	LogFlag   int
}

func defaultOptions(opt []Option) (option Option) {
	if len(opt) > 0 {
		option = opt[0]
	}

	if option.LogPrefix == "" {
		option.LogPrefix = "[default] "
	}

	return
}

func yayForOptions(opt ...Option) {
	option := defaultOptions(opt)

	log.SetPrefix(option.LogPrefix)
	log.SetFlags(option.LogFlag)

	log.Println("Hello, World!")
}
