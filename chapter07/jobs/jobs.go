package main

import (
	"fmt"
)

func main() {
	jobSample()
}

type Job struct {
	Value int
}

func jobSample() {
	joblist := []Job{{1}, {2}, {3}, {4}, {5}, {6}}

	jobs := make(chan Job)
	done := make(chan bool, len(joblist))

	go func() {
		for _, job := range joblist {
			jobs <- job // blocks waiting for a receive, because channel is unbuffered (size 0)
		}
		close(jobs)
	}()

	go func() {
		for job := range jobs { // blocks waiting for a send
			fmt.Printf("Job: %d\n", job.Value) // "do" the job
			done <- true
		}
	}()

	for i := 0; i < len(joblist); i++ {
		<-done // blocks waiting for a receive
	}
}
