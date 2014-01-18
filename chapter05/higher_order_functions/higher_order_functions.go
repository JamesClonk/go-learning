//#!/usr/bin/env goplay

package main

import (
	"fmt"
	"reflect"
)

func main() {
	letters := []string{"A", "B", "C", "D", "E", "F"}
	fmt.Printf("C index: %d\n", IndexReflect(letters, "C")) // huh? it doesn't work?
	fmt.Printf("4 index: %d\n", SliceIndex(len(letters), func(i int) bool { return letters[i] == "C" }))

	numbers := []int{1, 2, 3, 4, 5}
	fmt.Printf("4 index: %d\n", IndexReflect(numbers, 4)) // huh? it doesn't work?
	fmt.Printf("4 index: %d\n", SliceIndex(len(numbers), func(i int) bool { return numbers[i] == 4 }))
}

func IndexReflect(xs interface{}, x interface{}) int {
	if slice := reflect.ValueOf(xs); slice.Kind() == reflect.Slice {
		for i := 0; i < slice.Len(); i++ {
			if reflect.DeepEqual(x, slice.Index(i)) {
				return i
			}
		}
	}
	return -1
}

func SliceIndex(limit int, predicate func(int) bool) int {
	for i := 0; i < limit; i++ {
		if predicate(i) {
			return i
		}
	}
	return -1
}
