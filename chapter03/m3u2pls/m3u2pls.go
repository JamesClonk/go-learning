#!/usr/bin/env goplay

package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
)

var (
	extinfLineM3u = regexp.MustCompile(`\s*#EXTINF:([-]?[0-9]+),(.*)`)
	fileLineM3u   = regexp.MustCompile(`.*\.[[:word:]]{1,5}`)
)

type Track struct {
	Title  string
	Length int
	File   string
}

func (track *Track) IsComplete() bool {
	return track.Title != "" && track.Length != 0 && track.File != ""
}

func (track *Track) Reset() {
	track.Title = ""
	track.Length = 0
	track.File = ""
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("m3u2pls.go needs 1 argument: <m3u-file>")
		os.Exit(1)
	}

	ConvertM3uAndOutputPls(os.Args[1], os.Stdout)
}

func ConvertM3uAndOutputPls(m3ufilename string, output *os.File) {
	m3ufile, err := os.Open(m3ufilename)
	if err != nil {
		log.Fatal(err)
	}
	defer m3ufile.Close()

	reader := bufio.NewReader(m3ufile)
	writer := bufio.NewWriter(output)

	track := Track{"", 0, ""}
	fileCount := 0
	eof := false
	for !eof {
		line, err := reader.ReadString('\n')
		if err == io.EOF {
			eof = true
		} else if err != nil {
			log.Fatal(err)
		}

		if fileCount == 0 {
			if !strings.HasPrefix(line, "#EXTM3U") { // #EXTM3U
				log.Fatalf("Not a valid M3U file!")
			}
			if _, err := writer.WriteString("[playlist]\n"); err != nil {
				log.Fatal(err)
			}
			fileCount++

		} else if extinfLineM3u.MatchString(line) { // #EXTINF
			track.Reset()
			matches := extinfLineM3u.FindStringSubmatch(line)
			length, err := strconv.Atoi(matches[1])
			if err != nil {
				length = -1
			}
			track.Length = length
			track.Title = strings.Trim(matches[2], "\t ")

		} else if fileLineM3u.MatchString(line) { // File
			track.File = strings.Map(mapPlatformDirSeparator, strings.Trim(line, "\n\r\t "))
		}

		if track.IsComplete() { // -> pls
			writeTrack(writer, &track, &fileCount, "pls")
		}
	}

	footer := fmt.Sprintf("NumberOfEntries=%d\nVersion=%d\n", fileCount-1, 2)
	if _, err := writer.WriteString(footer); err != nil {
		log.Fatal(err)
	}
	writer.Flush()
}

func ConvertPlsAndOutputM3u(plsfilename string, output *os.File) {
	plsfile, err := os.Open(plsfilename)
	if err != nil {
		log.Fatal(err)
	}
	defer plsfile.Close()

	reader := bufio.NewReader(plsfile)
	writer := bufio.NewWriter(output)

	track := Track{"", 0, ""}
	fileCount := 0
	eof := false
	for !eof {
		line, err := reader.ReadString('\n')
		if err == io.EOF {
			eof = true
		} else if err != nil {
			log.Fatal(err)
		}

		if fileCount == 0 {
			if !strings.HasPrefix(line, "[playlist]") { // [playlist]
				log.Fatalf("Not a valid PLS file!")
			}
			if _, err := writer.WriteString("#EXTM3U\n"); err != nil {
				log.Fatal(err)
			}
			fileCount++

		} else if strings.HasPrefix(line, "File") { // File%d=...
			track.Reset()
			track.File = strings.Trim(line[strings.Index(line, "=")+1:], "\n\r\t ")

		} else if strings.HasPrefix(line, "Title") { // File%d=...
			track.Title = strings.Trim(line[strings.Index(line, "=")+1:], "\n\r\t ")

		} else if strings.HasPrefix(line, "Length") { // Length%d=...
			length, err := strconv.Atoi(strings.Trim(line[strings.Index(line, "=")+1:], "\n\r\t "))
			if err != nil {
				length = -1
			}
			track.Length = length
		}

		if track.IsComplete() { // -> m3u
			writeTrack(writer, &track, &fileCount, "m3u")
		}
	}
}

func writeTrack(writer *bufio.Writer, track *Track, fileCount *int, format string) {
	var data string
	if format == "m3u" {
		data = fmt.Sprintf("#EXTINF:%d,%s\n%s\n", track.Length, track.Title, track.File)
	} else {
		data = fmt.Sprintf("File%d=%s\nTitle%d=%s\nLength%d=%d\n", *fileCount, track.File, *fileCount, track.Title, *fileCount, track.Length)
	}

	if _, err := writer.WriteString(data); err != nil {
		log.Fatal(err)
	}
	writer.Flush()

	track.Reset()
	*fileCount++
}

// shamelessly copied this function
func mapPlatformDirSeparator(char rune) rune {
	if char == '/' || char == '\\' {
		return filepath.Separator
	}
	return char
}
