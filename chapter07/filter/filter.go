package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"runtime"
	"strings"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	minSize, maxSize, suffixes, files := handleCommandLine()
	sink(filterSize(minSize, maxSize, filterSuffixes(suffixes, source(files))))
}

func handleCommandLine() (minSize, maxSize int64, suffixes, files []string) {
	flag.Int64Var(&minSize, "min", -1, "minimum file size (-1 means no minimum)")
	flag.Int64Var(&maxSize, "max", -1, "maximum file size (-1 means no maximum)")
	var suffixesOpt *string = flag.String("suffixes", "", "comma-separated list of file suffixes")
	flag.Parse()

	if minSize > maxSize && maxSize != -1 {
		log.Fatalln("minimum size must be < maximum size")
	}

	suffixes = []string{}
	if *suffixesOpt != "" {
		suffixes = strings.Split(*suffixesOpt, ",")
	}

	files = flag.Args()

	return minSize, maxSize, suffixes, files
}

func source(files []string) <-chan string {
	filechannel := make(chan string, 100)

	go func() {
		for _, file := range files {
			filechannel <- file // non-blocking
		}
		close(filechannel) // sender closes channel
	}()

	return filechannel
}

func filterSuffixes(suffixes []string, in <-chan string) <-chan string {
	out := make(chan string, 100)

	go func() {
		for item := range in {
			for _, suffix := range suffixes {
				if strings.HasSuffix(item, suffix) {
					out <- item
				}
			}
		}
		close(out) // sender closes channel
	}()

	return out
}

func filterSize(min int64, max int64, in <-chan string) <-chan string {
	out := make(chan string, 100)

	go func() {
		for item := range in {
			file, err := os.Open(item)
			if err != nil {
				log.Printf("Could not open file [%s]: %v\n", item, err)
				continue
			}
			info, err := file.Stat()
			if err != nil {
				log.Println(err)
				continue
			}
			if info.Size() >= min && info.Size() <= max {
				out <- item
			}
		}
		close(out) // sender closes channel
	}()

	return out
}

func sink(in <-chan string) {
	for item := range in {
		fmt.Println(item)
	}
}
