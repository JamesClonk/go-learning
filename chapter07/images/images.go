package main

import (
	"fmt"
	"image"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"strings"
)

var workers = 5

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	if len(os.Args) < 2 || os.Args[1] == "-h" || os.Args[1] == "--help" {
		fmt.Printf("usage: %s <image files>\n", filepath.Base(os.Args[0]))
		os.Exit(1)
	}

	work := make(chan string, workers*2)
	done := make(chan struct{})

	go adder(work, commandLineFiles(os.Args[1:]))
	for i := 0; i < workers; i++ {
		go worker(done, work)
	}

	for i := 0; i < workers; i++ {
		<-done
	}
	close(done)
}

func adder(work chan string, files []string) {
	for _, filename := range files {
		work <- filename
	}
	close(work)
}

func worker(done chan struct{}, work chan string) {
	for filename := range work {
		process(filename)
	}
	done <- struct{}{}
}

func process(filename string) {
	filename = strings.Trim(filename, "\n\r\t ")

	file, err := os.Open(filename)
	if err != nil {
		log.Println("could not open file: ", err)
		return
	}
	defer file.Close()

	config, _, err := image.DecodeConfig(file)
	if err != nil {
		return
	}

	fmt.Printf(`<img src="%s" width="%d" height="%d" />`, filepath.Base(filename), config.Width, config.Height)
	fmt.Println()
}

func commandLineFiles(files []string) []string {
	if runtime.GOOS == "windows" {
		args := make([]string, 0, len(files))
		for _, name := range files {
			if matches, err := filepath.Glob(name); err != nil {
				args = append(args, name) // Invalid pattern
			} else if matches != nil { // At least one match
				args = append(args, matches...)
			}
		}
		return args
	}
	return files
}
