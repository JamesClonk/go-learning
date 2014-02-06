package main

import (
	"./safemap"
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"runtime"
)

var workers = runtime.NumCPU()

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	if len(os.Args) != 2 || os.Args[1] == "-h" || os.Args[1] == "--help" {
		fmt.Printf("usage: %s <report.log>\n", filepath.Base(os.Args[0]))
		os.Exit(1)
	}

	lines := make(chan string, workers*3)
	done := make(chan struct{}, workers)
	safeMap := safemap.New()
	go readLines(os.Args[1], lines)
	processLines(done, safeMap, lines)
	waitUntil(done)
	showResults(safeMap)
}

func readLines(filename string, lines chan<- string) {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal("could not open file:", err)
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	for {
		line, err := reader.ReadString('\n')
		if line != "" {
			lines <- line
		}
		if err != nil {
			if err != io.EOF {
				log.Println(err)
			}
			break
		}
	}
	close(lines)
}

func processLines(done chan<- struct{}, safeMap safemap.SafeMap, lines <-chan string) {
	getRx := regexp.MustCompile(`GET[ \t]+([^ \t\n]+[.]html?)`)
	incrementer := func(value interface{}, found bool) interface{} {
		if found {
			return value.(int) + 1
		}
		return 1
	}
	for i := 0; i < workers; i++ {
		go func() {
			for line := range lines {
				if matches := getRx.FindStringSubmatch(line); matches != nil {
					safeMap.Update(matches[1], incrementer)
				}
			}
			done <- struct{}{}
		}()
	}
}

func waitUntil(done <-chan struct{}) {
	for i := 0; i < workers; i++ {
		<-done
	}
}

func showResults(safeMap safemap.SafeMap) {
	pages := safeMap.Close()
	for page, count := range pages {
		fmt.Printf("%10d %s\n", count, page)
	}
}
