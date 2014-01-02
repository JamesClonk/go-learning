#!/usr/bin/env goplay

package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	arrays()
	slices()
	iterateOverSlices()
	modifyingSlices()
	sortandSearchSlices()
}

// arrays are fixed length
func arrays() {
	a := [5]int{1, 2, 3, 4, 5}
	fmt.Println(a)
}

// otherwise it's a slice (and slices are actually references to hidden arrays)
func slices() {
	a := make([]int, 5, 10)
	b := make([]int, 5)
	c := []int{}
	d := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(a, b, c, d)

	s := []string{"A", "B", "C", "D", "E"}
	t := s[:3]
	u := s[3:len(s)] // same as s[3:]
	fmt.Println(s, t, u)
	u[1] = "X"
	fmt.Println(s, t, u) // that's because slices are references.. don't forget this!

	// this is what happens behind the scenes (slices are references to hidden arrays)
	x := new([3]string)[:]
	x[0], x[1], x[2] = "A", "B", "C"
	fmt.Println(x)

	buffer := make([]int, 5, 10)
	buffer[0], buffer[2] = 1, 3
	fmt.Printf("%T %v %v %v\n", buffer, buffer, len(buffer), cap(buffer))
}

func iterateOverSlices() {
	s := []string{"A", "B", "C", "D", "E"}

	// copy element
	for _, e := range s[2:] {
		fmt.Println(e)
		e = "X"
	}
	fmt.Println(s)

	// actually modifying the slice
	for i, _ := range s {
		if i%2 == 0 {
			s[i] = "X"
		}
	}
	fmt.Println(s)
}

func modifyingSlices() {
	s := []string{"A", "B", "C", "D", "E"}
	fmt.Println(s, len(s), cap(s))

	s = append(s, "F", "G")
	fmt.Println(s, len(s), cap(s))

	s = append(s, s[:5]...)
	fmt.Println(s, len(s), cap(s))

	a := []string{"A", "B", "C", "D", "E"}
	b := []string{"X", "Y"}
	index := 2
	x := append(a[:index], append(b, a[index:]...)...) // insert slice 'b' into slice 'a' at position 'index'
	fmt.Println(a, b, x)

	x = append(x[:index], x[index+len(b):]...) // remove items in the middle of slice
	fmt.Println(x)
}

func sortandSearchSlices() {
	s := []string{"Alpha", "beta", "Charlie", "delta", "charly"}
	fmt.Println(s)

	sort.Strings(s) // case sensitive
	fmt.Println(s)

	sort.Sort(CaseInsenstiveSortableStrings(s)) // promote []string to CaseInsenstiveSortableStrings (which is also []string)
	fmt.Println(s)

	sort.Strings(s)
	index := sort.Search(len(s), func(i int) bool { return s[i] == "delta" }) // binary search, search for index of "delta"
	if index < len(s) && s[index] == "delta" {
		fmt.Printf("We really found it! At position %d\n", index)
	}
}

type CaseInsenstiveSortableStrings []string

func (slice CaseInsenstiveSortableStrings) Len() int {
	return len(slice)
}

func (slice CaseInsenstiveSortableStrings) Less(a, b int) bool {
	return strings.ToLower(slice[a]) < strings.ToLower(slice[b])
}

func (slice CaseInsenstiveSortableStrings) Swap(a, b int) {
	slice[a], slice[b] = slice[b], slice[a]
}
