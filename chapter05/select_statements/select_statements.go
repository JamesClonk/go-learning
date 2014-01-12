//#!/usr/bin/env goplay

package main

import (
	"fmt"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	first()
}

func first() {
	channels := make([]chan bool, 5)

	for i := range channels { // create 5 blocking channels
		channels[i] = make(chan bool)
	}

	go func() {
		for {
			// put value into one of the 5 channels at random
			channels[rand.Intn(5)] <- true
		}
	}()

	for i := 0; i < 30; i++ { // select first ready channel to read from, 30x times
		var x int
		select {
		case <-channels[0]:
			x = 1
		case <-channels[1]:
			x = 2
		case <-channels[2]:
			x = 3
		case <-channels[3]:
			x = 4
		case <-channels[4]:
			x = 5
		}
		fmt.Printf("%d ", x)
	}
	fmt.Println()
}
