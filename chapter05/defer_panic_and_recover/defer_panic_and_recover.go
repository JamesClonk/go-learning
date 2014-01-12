//#!/usr/bin/env goplay

package main

import (
	"fmt"
	"log"
)

func main() {
	defer func() { // catch-all recovery
		if r := recover(); r != nil {
			log.Printf("PANIC: %v\n", r)
		}
	}()

	deferSample1()

	x := deferSample2()
	fmt.Println(x)

	y := deferSample3()
	fmt.Println(*y)

	if err := panicAndRecovery(); err != nil {
		log.Println(err)
	}

	panicUp()

	fmt.Println("This here should never be seen..")
}

func deferSample1() {
	defer func() {
		fmt.Println("Hello! [from defer]")
	}()

	fmt.Println("World!")
}

func deferSample2() int {
	a := 1

	defer func() {
		a = 5
	}()

	return a
}

// dude, don't use crazy stuff like this in your code..
func deferSample3() *int {
	b := 1

	defer func() {
		b = 5
	}()

	return &b
}

func panicAndRecovery() (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("%v stopped..", r)
		}
	}()

	panic("Meltdown") // returns the function (works with named return parameters as seen here with "err")

	fmt.Println("This here should never be seen..")
	return nil
}

func panicUp() {
	panic("Oops!")
}
