#!/usr/bin/env goplay

package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"runtime"
)

type polar struct {
	radius float64
	degree float64
}

type cartesian struct {
	x float64
	y float64
}

var prompt = "Enter radius and angle in degrees, or %s to quit."

const result = "Polar radius=%.02f degree=%.02f -> Cartesian x=%.02f y=%.02f\n"

func init() {
	if runtime.GOOS == "windows" {
		prompt = fmt.Sprintf(prompt, "ctrl-z, enter")
	} else {
		prompt = fmt.Sprintf(prompt, "ctrl-c")
	}
}

func main() {
	questions := make(chan polar)
	defer close(questions)

	answers := solver(questions) // chan cartesian
	defer close(answers)

	interact(questions, answers)
}

func solver(questions chan polar) chan cartesian {
	answers := make(chan cartesian)

	// goroutine
	go func() {
		for {
			polarCoordinates := <-questions
			degree := polarCoordinates.degree * math.Pi / 180
			x := polarCoordinates.radius * math.Cos(degree)
			y := polarCoordinates.radius * math.Sin(degree)
			answers <- cartesian{x, y}
		}
	}()

	return answers
}

func interact(questions chan polar, answers chan cartesian) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println(prompt)

	for {
		fmt.Println("Radius and Degree: ")
		line, err := reader.ReadString('\n')
		if err != nil {
			break
		}

		var input polar
		if _, err := fmt.Sscanf(line, "%f %f", &input.radius, &input.degree); err != nil {
			fmt.Fprintln(os.Stderr, "invalid input")
			continue
		}

		// if _, err := fmt.Scanf("%f %f", &input.radius, &input.degree); err != nil {
		// 	fmt.Fprintln(os.Stderr, "invalid input")
		// 	continue
		// }

		questions <- input
		output := <-answers
		fmt.Printf(result, input.radius, input.degree, output.x, output.y)
	}

	fmt.Println()
}
