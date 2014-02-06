package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"runtime"
	"sync"
)

var workers = runtime.NumCPU()

type safeMap struct {
	data  map[string]int
	mutex *sync.RWMutex
}

func New() *safeMap {
	return &safeMap{make(map[string]int), new(sync.RWMutex)}
}

func (sm *safeMap) Increment(page string) {
	sm.mutex.Lock()
	defer sm.mutex.Unlock()
	sm.data[page]++
}

func (sm *safeMap) Len() int {
	sm.mutex.RLock()
	defer sm.mutex.RUnlock()
	return len(sm.data)
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	if len(os.Args) != 2 || os.Args[1] == "-h" || os.Args[1] == "--help" {
		fmt.Printf("usage: %s <report.log>\n", filepath.Base(os.Args[0]))
		os.Exit(1)
	}

	lines := make(chan string, workers*4)
	done := make(chan struct{}, workers)
	safeMap := New()

	go readLines(os.Args[1], lines)
	getRx := regexp.MustCompile(`GET[ \t]+([^ \t\n]+[.]html?)`)
	for i := 0; i < workers; i++ {
		go processLines(done, getRx, safeMap, lines)
	}

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

func processLines(done chan<- struct{}, getRx *regexp.Regexp, safeMap *safeMap, lines <-chan string) {
	for line := range lines {
		if matches := getRx.FindStringSubmatch(line); matches != nil {
			safeMap.Increment(matches[1])
		}
	}
	done <- struct{}{}
}

func waitUntil(done <-chan struct{}) {
	for i := 0; i < workers; i++ {
		<-done
	}
}

func showResults(safeMap *safeMap) {
	for page, count := range safeMap.data {
		fmt.Printf("%10d %s\n", count, page)
	}
}
