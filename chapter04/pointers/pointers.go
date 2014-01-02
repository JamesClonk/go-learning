#!/usr/bin/env goplay

package main

import (
	"fmt"
)

func main() {
	example1()
	example2()
	example3()
}

func example1() {
	i := 9
	j := 5
	var product int

	swapAndProduct := func(a, b, p *int) {
		*a, *b = *b, *a
		*p = *a * *b
	}

	swapAndProduct(&i, &j, &product)
	fmt.Println(i, j, product)
}

func example2() {
	type composer struct {
		name      string
		birthyear int
	}

	antónio := composer{"António Teixeira", 1707} // composer value

	agnes := new(composer) // pointer to empty(default values) composer
	agnes.name, agnes.birthyear = "Agnes Zimmermann", 1845

	julia := &composer{} // pointer to empty(default values) composer
	julia.name, julia.birthyear = "Julia Ward Howe", 1819

	augusta := &composer{"August Holmès", 1847} // pointer to composer

	fmt.Println(antónio)
	fmt.Println(agnes, augusta, julia)
}

// slices are references
func example3() {
	nums := []int{1, 2, 3, 4, 5}

	increase := func(numbers []int, inc int) {
		for i, _ := range numbers { // for i := range numbers
			numbers[i] += inc
		}
	}

	increase(nums, 2)
	fmt.Println(nums)
}
