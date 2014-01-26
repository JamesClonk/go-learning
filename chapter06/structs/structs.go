package main

import (
	"fmt"
)

func main() {
	anonStructSample()
	embeddingSample()
	optionSample()
}

func anonStructSample() {
	// slice of anonymous struct with 7 numbers
	gateCoordinates := []struct{ x, y, z, a, b, c, o int }{{1, 2, 3, 4, 5, 6, 7}, {}, {3, 4, 2, 6, 5, 7, 0}}
	for _, chevron := range gateCoordinates {
		fmt.Printf("(%d,%d,%d,%d,%d,%d,%d)\n", chevron.x, chevron.y, chevron.z, chevron.a, chevron.b, chevron.c, chevron.o)
	}
}

type Person struct {
	FirstName string
	LastName  string
}

type Author struct {
	Person   // embedded
	YearBorn int
}

func (p Person) Name() string {
	return fmt.Sprintf("%s, %s", p.LastName, p.FirstName)
}

func embeddingSample() {
	author := Author{Person{"Mr.", "Smart"}, 1900}
	fmt.Println(author)

	author.FirstName = "Ms."
	author.YearBorn += 15
	author.Person.LastName = "Smarter"
	fmt.Println(author)

	fmt.Println(author.Name())
}

type Optioner interface {
	Name() string
	IsValid() bool
}

type OptionCommon struct {
	ShortName string "short option name"
	LongName  string "long option name"
}

type IntOption struct {
	OptionCommon    // embedded
	Value, Min, Max int
}

func (option IntOption) Name() string {
	return name(option.ShortName, option.LongName)
}

func (option IntOption) IsValid() bool {
	return option.Min <= option.Value && option.Value <= option.Max
}

func name(short, long string) string {
	if long == "" {
		return short
	}
	return long
}

func optionSample() {
	numOpt := IntOption{
		OptionCommon: OptionCommon{"n", "number"},
		Min:          1,
		Max:          10,
	}
	fmt.Printf("name=%s\nmin=%d\n", numOpt.Name(), numOpt.Min)
}
