#!/usr/bin/env goplay

package main

import (
	"fmt"
	"regexp"
	"sort"
	"strings"
)

func main() {
	fmt.Println(UniqueInts([]int{1, 2, 5, 6, 7, 8, 9, 0, 5, 5, 5, 5, 5, 5, 5, 7, 2, 0, 2}))
	fmt.Println(Flatten([][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}))
	fmt.Println(Make2D([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}, 3))
	fmt.Println(Make2D([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}, 4))
	fmt.Println(Make2D([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}, 6))

	iniData := []string{
		"; Cut down copy of Mozilla application.ini file",
		"",
		"Special=Test",
		"[App]",
		"Vendor=Mozilla",
		"Name=Iceweasel",
		"Profile=mozilla/firefox",
		"Version=3.5.16",
		"[Gecko]",
		"MinVersion=1.9.1",
		"MaxVersion=1.9.1.*",
		"[XRE]",
		"EnableProfileMigrator=0",
		"EnableExtensionManager=1",
	}
	ini := ParseIni(iniData)
	fmt.Println(ini)
	PrintIni(ini)
}

// Chapter 4.5 - Exercise 1
func UniqueInts(input []int) (output []int) {
	mem := map[int]int{}
	for _, value := range input {
		if _, found := mem[value]; !found {
			mem[value] = value
			output = append(output, value)
		}
	}
	return

	/*
		or:
			for _, value := range input {
				mem[value] = value
			}
			for key, _ := range mem {
				output = append(output, key)
			}
	*/
}

// Chapter 4.5 - Exercise 2
func Flatten(input [][]int) (output []int) {
	for _, row := range input {
		for _, value := range row {
			output = append(output, value)
		}
	}
	return
}

// Chapter 4.5 - Exercise 3
func Make2D(input []int, columns int) (result [][]int) {
	r, c := -1, 0
	for _, value := range input {
		if c == 0 {
			result = append(result, make([]int, columns))
			r++
		}
		result[r][c] = value

		c++
		if c%columns == 0 {
			c = 0
		}
	}
	return
}

// Chapter 4.5 - Exercise 4
func ParseIni(input []string) map[string]map[string]string {
	result := map[string]map[string]string{}
	groupRx := regexp.MustCompile(`\[([[:word:]]+)\]`)
	settingRx := regexp.MustCompile(`([[:word:]]+)=(.*)`)
	currentGroup := "General"
	for _, line := range input {
		if groupRx.MatchString(line) {
			currentGroup = strings.Trim(groupRx.FindStringSubmatch(line)[1], "\t\n\r ")
		} else if settingRx.MatchString(line) {
			matched := settingRx.FindStringSubmatch(line)
			if _, found := result[currentGroup]; !found {
				result[currentGroup] = map[string]string{}
			}
			result[currentGroup][matched[1]] = strings.Trim(matched[2], "\t\n\r ")
		}
	}
	return result
}

// Chapter 4.5 - Exercise 5
func PrintIni(input map[string]map[string]string) {
	groups := make([]string, 0, len(input))
	for group, _ := range input {
		groups = append(groups, group)
	}
	sort.Strings(groups)

	for _, group := range groups {
		keys := []string{}
		for key, _ := range input[group] {
			keys = append(keys, key)
		}
		sort.Strings(keys)

		fmt.Printf("[%s]\n", group)
		for _, key := range keys {
			fmt.Printf("%s=%s\n", key, input[group][key])
		}
		fmt.Print("\n")
	}
}
