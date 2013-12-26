#!/usr/bin/env goplay

package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

var bigDigits = [][]string{{
	"  000  ",
	" 0   0 ",
	"0     0",
	"0     0",
	"0     0",
	" 0   0 ",
	"  000  "}, {
	"  1  ",
	" 11  ",
	"  1  ",
	"  1  ",
	"  1  ",
	"  1  ",
	" 111 "}, {
	" 222 ",
	"2   2",
	"   2 ",
	"  2  ",
	" 2   ",
	"2    ",
	"22222"}, {
	" 333 ",
	"3   3",
	"    3",
	"  33 ",
	"    3",
	"3   3",
	" 333 "}, {
	"   4  ",
	"  44  ",
	" 4 4  ",
	"4  4  ",
	"444444",
	"   4  ",
	"   4  "}, {
	"55555",
	"5    ",
	"5    ",
	" 555 ",
	"    5",
	"5   5",
	" 555 "}, {
	" 666 ",
	"6    ",
	"6    ",
	"6666 ",
	"6   6",
	"6   6",
	" 666 "}, {
	"77777",
	"    7",
	"   7 ",
	"  7  ",
	" 7   ",
	"7    ",
	"7    "}, {
	" 888 ",
	"8   8",
	"8   8",
	" 888 ",
	"8   8",
	"8   8",
	" 888 "}, {
	" 9999",
	"9   9",
	"9   9",
	" 9999",
	"    9",
	"    9",
	"    9"},
}

func main() {
	if len(os.Args) < 2 {
		usage()
	}

	for _, arg := range os.Args[1:] {
		if arg == "-h" || arg == "--help" {
			usage()
		}
	}

	var bar bool // default is false
	for _, arg := range os.Args[1:] {
		if arg == "-b" || arg == "--bar" {
			bar = true
		}
	}

	input := os.Args[1]
	log.Println("Input: " + input)

	var lines string
	for row := 0; row < len(bigDigits[0]); row++ {
		for index := range input {
			digit := input[index] - '0' // ascii trickery
			lines += bigDigits[digit][row]
		}
		lines += "\n"
	}

	if bar {
		lines = addBars(lines)
	}

	fmt.Print(lines)
}

func usage() {
	fmt.Printf("usage: %s [-b|--bar] <whole-number>\n"+
		"-b --bar   draw an over- and underbar\n", filepath.Base(os.Args[0]))
	os.Exit(1)
}

func addBars(lines string) string {
	bar := fmt.Sprint(strings.Repeat("*", len(strings.Split(lines, "\n")[0]))) + "\n"
	return bar + lines + bar
}
