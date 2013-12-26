package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

func main() {
	inFilename, outFilename, err := parseCommandline()
	if err != nil {
		//log.Fatal(err)
		fmt.Println(err)
		os.Exit(1)
	}

	inFile, outFile := os.Stdin, os.Stdout // default in/out

	// overwrite in?
	if inFilename != nil {
		if inFile, err = os.Open(*inFilename); err != nil {
			log.Fatal(err)
		}
		defer inFile.Close()
	}

	// overwrite out?
	if outFilename != nil {
		if outFile, err = os.Create(*outFilename); err != nil {
			log.Fatal(err)
		}
		defer outFile.Close()
	}

	if err = americanise(inFile, outFile); err != nil {
		log.Fatal(err)
	}
}

func parseCommandline() (inFilename, outFilename *string, err error) {
	if len(os.Args) > 1 {
		if os.Args[1] == "-h" || os.Args[1] == "--help" {
			return nil, nil, fmt.Errorf("usage: %s <input.txt> <output.txt>", filepath.Base(os.Args[0]))
		}

		inFilename = &os.Args[1]

		if len(os.Args) > 2 {
			outFilename = &os.Args[2]
		}
	}

	if inFilename != nil && inFilename == outFilename {
		log.Fatal("input and output filenames cannot be the same!")
	}

	// bare 'return' would also be possible
	return inFilename, outFilename, nil
}

func americanise(inFile io.Reader, outFile io.Writer) (err error) {
	reader := bufio.NewReader(inFile)
	writer := bufio.NewWriter(outFile)

	// hmmm.. nice!
	defer func() {
		if err == nil {
			err = writer.Flush()
		}
	}() // call

	var replacer func(string) string
	if replacer, err = replacerFunction(); err != nil {
		return err
	}

	rgx := regexp.MustCompile("[A-Za-z]+")
	eof := false
	for !eof {
		var line string
		line, err = reader.ReadString('\n')
		if err == io.EOF {
			err = nil // I do not like this io.EOF 'error'.. oh well..
			eof = true
		} else if err != nil {
			return err
		}

		line = rgx.ReplaceAllStringFunc(line, replacer)
		if _, err = writer.WriteString(line); err != nil {
			return err
		}
	}

	return err
}

func replacerFunction() (func(string) string, error) {
	bytes, err := ioutil.ReadFile("british-american.txt")
	if err != nil {
		return nil, err
	}

	text := string(bytes)

	// setup mapping
	usForBritish := make(map[string]string)
	lines := strings.Split(text, "\n")
	for _, line := range lines {
		fields := strings.Fields(line)
		if len(fields) == 2 {
			usForBritish[fields[0]] = fields[1]
		}
	}

	replacer := func(input string) string {
		if us, found := usForBritish[input]; found {
			return us
		}
		return input
	}

	return replacer, nil
}
