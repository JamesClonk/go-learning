//#!/usr/bin/env goplay

package main

import (
	"fmt"
)

func main() {
	sample()
	doneExample() // this show how to use a "done" channel, to check if a goroutine is done with work and block until then..

	sequence := createSequence(100, 5)
	for i := 0; i < 20; i++ {
		x := <-sequence
		fmt.Printf("(X -> %d), (X2 -> %d)\n", x, <-sequence)
	}
}

func sample() {
	c := make(chan int)

	a := 0
	go func() { // run this in another goroutine
		for {
			if a >= 10 {
				break
			}
			a++
			c <- a
		}
	}()

	fmt.Println(a) // unreliable.. ofc!
	for {
		x := <-c
		fmt.Println(a, x)
		if x >= 10 {
			break
		}
	}
}

func createSequence(start int, cache int) chan int {
	next := make(chan int, cache)

	go func(i int) {
		for {
			next <- i
			i++
		}
	}(start)

	return next
}

func doneExample() {
	done := make(chan bool)

	a := 0
	go func() {
		for {
			if a >= 10 {
				done <- true
				break
			}
			a++
		}
	}()

	select { // wait for channel 'done' to contain a message
	case <-done:
		fmt.Println("done!", a)
	}
}
