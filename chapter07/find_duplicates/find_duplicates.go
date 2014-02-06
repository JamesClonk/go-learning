package main

import (
	"crypto/sha1"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"sort"
	"sync"
)

const maxGoroutines = 100
const maxSizeOfSmallFile = 1024 * 32

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	if len(os.Args) == 1 || os.Args[1] == "-h" || os.Args[1] == "--help" {
		fmt.Printf("usage: %s <path>\n", filepath.Base(os.Args[0]))
		os.Exit(1)
	}

	info := make(chan fileInfo, maxGoroutines*2)
	go findDuplicates(info, os.Args[1])
	data := mergeResults(info)
	outputResults(data)
}

type fileInfo struct {
	sha1 []byte
	size int64
	path string
}

type pathsInfo struct {
	size  int64
	paths []string
}

func findDuplicates(info chan fileInfo, dirname string) {
	waiter := &sync.WaitGroup{}
	filepath.Walk(dirname, makeWalkFunc(info, waiter))
	waiter.Wait()
	close(info)
}

func makeWalkFunc(info chan fileInfo, waiter *sync.WaitGroup) func(string, os.FileInfo, error) error {
	return func(path string, fi os.FileInfo, err error) error {
		if err == nil && fi.Size() > 0 &&
			(fi.Mode()&os.ModeType == 0) {
			if fi.Size() < maxSizeOfSmallFile ||
				runtime.NumGoroutine() > maxGoroutines {
				processFile(path, fi, info, nil)
			} else {
				waiter.Add(1)
				go processFile(path, fi, info,
					func() { waiter.Done() })
			}
		}
		return nil
	}
}

func processFile(filename string, fi os.FileInfo, info chan fileInfo, done func()) {
	if done != nil {
		defer done()
	}

	file, err := os.Open(filename)
	if err != nil {
		log.Println(err)
		return
	}
	defer file.Close()

	hash := sha1.New()
	if size, err := io.Copy(hash, file); size != fi.Size() || err != nil {
		if err != nil {
			log.Println(err)
		} else {
			log.Println("could not read whole file:", filename)
		}
		return
	}

	info <- fileInfo{hash.Sum(nil), fi.Size(), filename}
}

func mergeResults(info <-chan fileInfo) map[string]*pathsInfo {
	data := make(map[string]*pathsInfo)
	format := fmt.Sprintf("%%016X:%%%dX", sha1.Size*2)

	for info := range info {
		key := fmt.Sprintf(format, info.size, info.sha1)
		value, found := data[key]
		if !found {
			value = &pathsInfo{size: info.size}
			data[key] = value
		}
		value.paths = append(value.paths, info.path)
	}

	return data
}

func outputResults(data map[string]*pathsInfo) {
	keys := make([]string, 0, len(data))
	for key := range data {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	for _, key := range keys {
		value := data[key]
		if len(value.paths) > 1 {
			fmt.Printf("%d duplicate files (%s bytes):\n",
				len(value.paths), commas(value.size))
			sort.Strings(value.paths)
			for _, name := range value.paths {
				fmt.Printf("\t%s\n", name)
			}
		}
	}
}

func commas(x int64) string {
	value := fmt.Sprint(x)
	for i := len(value) - 3; i > 0; i -= 3 {
		value = value[:i] + "," + value[i:]
	}
	return value
}
