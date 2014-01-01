#!/usr/bin/env goplay

package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"regexp"
	"strings"
)

func main() {
	bytes, err := ioutil.ReadFile("example.conf")
	if err != nil {
		log.Fatalf("Could not read file [%s]: %s", "example.conf", err)
	}

	config := make(map[string]string)
	rx := regexp.MustCompile(`\s*([[:alpha:]]\w*)\s+(.+)`)
	if matched := rx.FindAllStringSubmatch(string(bytes), -1); matched != nil {
		for _, match := range matched {
			config[match[1]] = strings.Trim(match[2],"\t ")
		}
	}

	fmt.Print(config)
}