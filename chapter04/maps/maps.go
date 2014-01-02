//#!/usr/bin/env goplay

package main

import (
	"fmt"
	"sort"
)

// maps are references too..
func main() {
	example1()
	example2()
}

func example1() {
	a := make(map[string]int, 10)
	b := make(map[string]int)
	c := map[string]int{}
	d := map[string]int{"a": 1, "b": 2, "c": 3, "d": 4, "e": 5}
	d["f"], d["g"] = 6, 7

	x := [4]map[string]int{a, b, c, d} // an array of maps (empty [] would make it a slice)
	for _, e := range x {
		fmt.Println(e, len(e))
	}

	// map lookup
	fmt.Println(d["f"], d["a"], d["x"]) // if key is not found, it will return the empty/default value of type, "x" = 0

	// with this lookup syntax we can check for key existance
	if key, found := d["b"]; found {
		fmt.Println(key)
	}
	if _, found := d["x"]; !found {
		fmt.Println("Not found!")
	}

	delete(d, "a") // delete
	delete(d, "y") // does nothing
	d["b"] = 999   // update
	d["x"] = 777   // insert
	fmt.Println(d)

	// let's order & print out the map
	keys := make([]string, 0, len(d))
	for key := range d {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	line := "map["
	for _, key := range keys {
		line += fmt.Sprintf("%v:%v", key, d[key])
		line += " "
	}
	fmt.Println(line[:len(line)-1] + "]")
}

type point struct {
	x int
	y int
}

func (p *point) String() string {
	return fmt.Sprintf("(%d,%d)", p.x, p.y)
}

// even pointers can be used as keys
func example2() {
	p := make(map[*point]string, 3) // always try to use initial capacity if possible, it helps performance
	p[&point{4, 5}] = "Alpha"
	p[&point{7, 6}] = "Beta"
	p[&point{9, -1}] = "Omega"

	fmt.Println(p)
	for pt, name := range p { // range over a map
		fmt.Println(pt, "=>", name)
	}

	// invert map
	px := make(map[string]*point, len(p))
	for key, value := range p {
		px[value] = key
	}
	fmt.Println(px)
}
