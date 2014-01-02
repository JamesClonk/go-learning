//#!/usr/bin/env goplay

package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Printf("usage: %s <inputfile>\n", filepath.Base(os.Args[0]))
		os.Exit(1)
	}

	separators := []string{"|", "\t", "*", ";"}

	lines := readLinesN(os.Args[1], 5)
	counts := createCounts(lines, separators)
	separator := guessSeparator(counts)
	fmt.Printf("Separator is likely to be: %q\n", separator)
}

func readLinesN(filename string, linesToRead int) (lines []string) {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	for i := 0; i < linesToRead; i++ {
		line, err := reader.ReadString('\n')
		if err == io.EOF {
			i = linesToRead
		} else if err != nil {
			log.Fatal(err)
		}
		lines = append(lines, line)
	}
	return
}

func createCounts(lines []string, separators []string) map[string]int {
	result := make(map[string]int, len(separators))
	for _, separator := range separators {
		for _, line := range lines {
			count := strings.Count(line, separator)
			result[separator] += count
		}
	}
	return result
}

func guessSeparator(counts map[string]int) (separator string) {
	var count int
	for key, value := range counts {
		if value > count {
			count = value
			separator = key
		}
	}
	return
}
