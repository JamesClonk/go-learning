#!/usr/bin/env goplay

package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strings"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Printf("usage: %s <filename>\n", filepath.Base(os.Args[0]))
		os.Exit(1)
	}

	data, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	lines := strings.SplitN(string(data), "\n", -1)

	rx := regexp.MustCompile("[[:word:]]+")
	frequencies := make(map[string]int, len(lines))
	for _, line := range lines {
		words := rx.FindAllString(strings.Trim(strings.ToLower(line), " \t\n\r"), -1)
		for _, word := range words {
			frequencies[word]++
		}
	}

	reportByWords(frequencies)
	reportByFrequency(frequencies)
}

func reportByWords(frequencies map[string]int) {
	words := make([]string, len(frequencies))
	i := 0
	for key, _ := range frequencies {
		words[i] = key
		i++
	}
	sort.Strings(words)
	for _, word := range words {
		fmt.Printf("%-20s %d\n", word, frequencies[word])
	}
}

func reportByFrequency(freqsByWord map[string]int) {
	wordsByFreq := make(map[int][]string)
	for word, freq := range freqsByWord {
		wordsByFreq[freq] = append(wordsByFreq[freq], word)
	}

	frequencies := make([]int, len(wordsByFreq))
	i := 0
	for freq, _ := range wordsByFreq {
		frequencies[i] = freq
		i++
	}
	sort.Ints(frequencies)
	for _, freq := range frequencies {
		fmt.Printf("%-4d %s\n", freq, wordsByFreq[freq])
	}
}
