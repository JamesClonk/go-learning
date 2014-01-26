package main

import (
	"fmt"
)

func main() {
	Swapping()
}

func Swapping() {
	jekyll := Tuple{"Henry", "Jekyll"}
	hyde := Tuple{"Edward", "Hyde"}
	fmt.Println("Before: ", jekyll, hyde)

	jekyll.Swap()
	hyde.Swap()
	fmt.Println("After #1: ", jekyll, hyde)

	swapStuff(&jekyll, &hyde)
	fmt.Println("After #2: ", jekyll, hyde)
}

type Swapper interface {
	Swap()
}

type Tuple struct{ _1, _2 string }

func (t *Tuple) Swap() {
	t._1, t._2 = t._2, t._1
}

func (t *Tuple) String() string {
	return fmt.Sprintf("(%s, %s)", t._1, t._2)
}

func swapStuff(stuff ...Swapper) {
	for i, _ := range stuff {
		stuff[i].Swap()
	}
}
