package main

import (
	"fmt"
	"os"
	"strconv"
)

type memoize func(int, ...int) interface{}

var Fibonacci memoize

func init() {
	Fibonacci = Memoize(func(x int, xs ...int) interface{} {
		if x < 2 {
			return x
		}
		return Fibonacci(x-1).(int) + Fibonacci(x-2).(int)
	})
}

func main() {
	if len(os.Args) == 2 {
		n, err := strconv.Atoi(os.Args[1])
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		fmt.Println(Fibonacci(n))
	}
}

func Memoize(f memoize) memoize {
	cache := make(map[string]interface{})
	return func(x int, xs ...int) interface{} {
		key := fmt.Sprint(x)
		for _, i := range xs {
			key += fmt.Sprintf(",%d", i)
		}

		if value, found := cache[key]; found {
			return value
		}

		value := f(x, xs...)
		cache[key] = value
		return value
	}
}
