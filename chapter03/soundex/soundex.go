//#!/usr/bin/env goplay

package main

import (
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
	"strings"
)

type soundex struct {
	Name    string
	Soundex string
}

type soundexTest struct {
	Name     string
	Soundex  string
	Expected string
	Test     string
}

var templates = template.Must(template.ParseGlob("*.html"))

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/test", test)
	if err := http.ListenAndServe(":9000", nil); err != nil {
		log.Fatal("http server failed to start", err)
	}
}

func index(response http.ResponseWriter, request *http.Request) {
	log.Println(*request)

	if err := request.ParseForm(); err != nil {
		http.Error(response, err.Error(), http.StatusInternalServerError)
		return
	}

	var sx []soundex
	if names, ok := parseRequest(request); ok {
		for _, name := range names {
			sx = append(sx, soundex{name, getSoundexValue(name)})
		}
	}

	if err := templates.ExecuteTemplate(response, "index.html", sx); err != nil {
		http.Error(response, err.Error(), http.StatusInternalServerError)
	}
}

func test(response http.ResponseWriter, request *http.Request) {
	log.Println(*request)

	var sx []soundexTest

	bytes, err := ioutil.ReadFile("soundex-test-data.txt")
	if err != nil {
		log.Fatal(err)
	}

	rx := regexp.MustCompile(`[A-Z][\d]+ [[:word:]]`)
	for _, line := range strings.SplitN(string(bytes), "\n", -1) {
		if rx.MatchString(line) {
			name := line[strings.Index(line, " ")+1:]
			expected := line[:strings.Index(line, " ")]
			value := getSoundexValue(name)
			test := "FAIL"
			if value == expected {
				test = "PASS"
			}
			sx = append(sx, soundexTest{name, value, expected, test})
		}
	}

	if err := templates.ExecuteTemplate(response, "test.html", sx); err != nil {
		http.Error(response, err.Error(), http.StatusInternalServerError)
	}
}

func parseRequest(request *http.Request) ([]string, bool) {
	var names []string
	if data, found := request.Form["names"]; found && len(data) > 0 {
		text := strings.Replace(data[0], ",", " ", -1)
		for _, name := range strings.Fields(text) {
			names = append(names, name)
		}
	}
	if len(names) == 0 {
		return names, false
	}
	return names, true
}

var digitForLetter = []rune{
	0, 1, 2, 3, 0, 1, 2, 0, 0, 2, 2, 4, 5,
	5, 0, 1, 2, 6, 2, 3, 0, 1, 0, 2, 0, 2}

// "c - 'A'" produces a 0-based index, so 'A' -> 0, 'B' -> 1, etc.
// "'0' + digitForLetter[index]" converts a one digit integer into the
// equivalent Unicode character, i.e., 0 -> "0", 1 -> "1", etc.
func getSoundexValue(name string) string {
	name = strings.ToUpper(name)
	chars := []rune(name)
	var codes []rune
	codes = append(codes, chars[0])
	for i := 1; i < len(chars); i++ {
		char := chars[i]
		if i > 0 && char == chars[i-1] {
			continue
		}
		if index := char - 'A'; 0 <= index &&
			index < int32(len(digitForLetter)) &&
			digitForLetter[index] != 0 {
			codes = append(codes, '0'+digitForLetter[index])
		}
	}
	for len(codes) < 4 {
		codes = append(codes, '0')
	}
	return string(codes[:4])
}
