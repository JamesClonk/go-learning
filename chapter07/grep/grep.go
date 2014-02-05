package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"runtime"
	"time"
)

type Job struct {
	filename string
	results  chan<- Result
}

type Result struct {
	filename string
	lino     int
	line     string
}

func (job Job) Do(rx *regexp.Regexp) {
	file, err := os.Open(job.filename)
	if err != nil {
		log.Printf("error: %s\n", err)
		return
	}
	defer file.Close()
	reader := bufio.NewReader(file)
	for lino := 1; ; lino++ {
		line, err := reader.ReadBytes('\n')
		line = bytes.TrimRight(line, "\n\r")
		if rx.Match(line) {
			job.results <- Result{job.filename, lino, string(line)}
		}
		if err != nil {
			if err != io.EOF {
				log.Printf("error:%d: %s\n", lino, err)
			}
			break
		}
	}
}

var workers = runtime.NumCPU()

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	if len(os.Args) < 3 || os.Args[1] == "-h" || os.Args[1] == "--help" {
		fmt.Printf("usage: %s <regexp> <files>\n", filepath.Base(os.Args[0]))
		os.Exit(1)
	}

	if rx, err := regexp.Compile(os.Args[1]); err != nil {
		log.Fatalf("invalid regexp: %s\n", err)
	} else {
		grep(rx, os.Args[2:])
	}
}

func grep(rx *regexp.Regexp, files []string) {
	jobs := make(chan Job, workers)
	results := make(chan Result, minimum(1000, len(files)))
	done := make(chan struct{}, workers)

	go addJobs(jobs, files, results)
	for i := 0; i < workers; i++ {
		go doJobs(done, rx, jobs)
	}
	go awaitCompletion(done, results) // not really necessary to have it's own goroutine

	processResults(results) // main goroutine / blocking
}

func addJobs(jobs chan Job, files []string, results chan Result) {
	for _, file := range files {
		jobs <- Job{file, results}
	}
	close(jobs)
}

func doJobs(done chan struct{}, rx *regexp.Regexp, jobs chan Job) {
	for job := range jobs {
		job.Do(rx)
	}
	done <- struct{}{}
}

func awaitCompletion(done chan struct{}, results chan Result) {
	for i := 0; i < workers; i++ {
		<-done
	}
	close(done)
	close(results)
}

func processResults(results chan Result) {
	for result := range results {
		fmt.Printf("%8.20s [%d]: %s\n", result.filename, result.lino, result.line)
	}
}

func minimum(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

// with timeout
func waitAndProcessResults(timeout int64, done <-chan struct{},
	results <-chan Result) {
	finish := time.After(time.Duration(timeout))
	for working := workers; working > 0; {
		select { // Blocking
		case result := <-results:
			fmt.Printf("%s:%d:%s\n", result.filename, result.lino,
				result.line)
		case <-finish:
			fmt.Println("timed out")
			return // Time's up so finish with what results there were
		case <-done:
			working--
		}
	}
	for {
		select { // Nonblocking
		case result := <-results:
			fmt.Printf("%s:%d:%s\n", result.filename, result.lino,
				result.line)
		case <-finish:
			fmt.Println("timed out")
			return // Time's up so finish with what results there were
		default:
			return
		}
	}
}
