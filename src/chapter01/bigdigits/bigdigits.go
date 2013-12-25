package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
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
	if len(os.Args) != 2 {
		fmt.Printf("usage: %s <number>\n", filepath.Base(os.Args[0]))
		os.Exit(1)
	}

	input := os.Args[1]
	log.Println("Input: " + input)

	for row := 0; row < len(bigDigits[0]); row++ {
		line := ""
		for index := range input {
			digit := input[index] - '0' // ascii trickery
			line += bigDigits[digit][row]
		}
		fmt.Println(line)
	}
}
