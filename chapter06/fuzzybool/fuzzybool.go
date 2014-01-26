package main

import (
	"fmt"
	"log"
)

func main() {
	f1, err := New(0.56)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(f1)

	if err := f1.Set(0.21); err != nil {
		log.Fatal(err)
	}
	fmt.Println(f1)
}

type FuzzyBool struct {
	value float64
}

func (f *FuzzyBool) String() string {
	return fmt.Sprintf("%.0f%%", 100*f.value)
}

func (f *FuzzyBool) Set(value interface{}) (err error) {
	f.value, err = float64Value(value)
	return err
}

func (f *FuzzyBool) Copy() *FuzzyBool {
	return &FuzzyBool{f.value}
}

func (f *FuzzyBool) Not() *FuzzyBool {
	return &FuzzyBool{1 - f.value}
}

func (f *FuzzyBool) And(first *FuzzyBool, rest ...*FuzzyBool) *FuzzyBool {
	min := f.value
	rest = append(rest, first)
	for _, other := range rest {
		if min > other.value {
			min = other.value
		}
	}
	return &FuzzyBool{min}
}

func (f *FuzzyBool) Or(first *FuzzyBool, rest ...*FuzzyBool) *FuzzyBool {
	min := f.value
	rest = append(rest, first)
	for _, other := range rest {
		if min < other.value {
			min = other.value
		}
	}
	return &FuzzyBool{min}
}

func (f *FuzzyBool) Less(other *FuzzyBool) bool {
	return f.value < other.value
}

func (f *FuzzyBool) Equal(other *FuzzyBool) bool {
	return f.value == other.value
}

func (f *FuzzyBool) Bool() bool {
	return f.value >= 0.5
}

func (f *FuzzyBool) Float() float64 {
	return f.value
}

func New(value interface{}) (*FuzzyBool, error) {
	val, err := float64Value(value)
	return &FuzzyBool{val}, err
}

func float64Value(value interface{}) (fuzzy float64, err error) {
	switch value := value.(type) {
	case float32:
		fuzzy = float64(value)
	case float64:
		fuzzy = value
	case int:
		fuzzy = float64(value)
	case bool:
		fuzzy = 0
		if value {
			fuzzy = 1
		}
	default:
		return 0, fmt.Errorf("float64ForValue(): %v is not a number or boolean", value)
	}

	if fuzzy < 0 {
		fuzzy = 0
	} else if fuzzy > 1 {
		fuzzy = 1
	}
	return fuzzy, nil
}
