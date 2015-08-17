package main

import (
	"bytes"
	"fmt"
	"strings"
)

// MemoizeFunction ...
type MemoizeFunction func(int, ...int) interface{}

// Fibonacci ...
var Fibonacci MemoizeFunction

// RomanForDecimal ...
var RomanForDecimal MemoizeFunction

func init() {
	Fibonacci = Memoize(func(x int, xs ...int) interface{} {
		if x < 2 {
			return x
		}
		return Fibonacci(x-1).(int) + Fibonacci(x-2).(int)
	})

	decimals := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	romans := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}
	RomanForDecimal = Memoize(func(x int, xs ...int) interface{} {
		if x < 0 || x > 3999 {
			panic("RomanForDecimal() only handles integers [0, 3999]")
		}

		var buffer bytes.Buffer
		for i, decimal := range decimals {
			remainder := x / decimal
			x %= decimal
			if remainder > 0 {
				_, _ = buffer.WriteString(strings.Repeat(romans[i], remainder))
			}
		}
		return buffer.String()
	})
}

func main() {
	fmt.Println("Fibonacci(45) = ", Fibonacci(45).(int))

	fmt.Println("RomanForDecimal(45) = ", RomanForDecimal(45).(string))
	fmt.Println("RomanForDecimal(2345) = ", RomanForDecimal(2345).(string))
}

// Memoize ...
func Memoize(function MemoizeFunction) MemoizeFunction {
	cache := make(map[string]interface{})

	return func(x int, xs ...int) interface{} {
		key := fmt.Sprint(x)
		for _, i := range xs {
			key += fmt.Sprintf(",%d", i)
		}
		if value, found := cache[key]; found {
			return value
		}
		value := function(x, xs...)
		cache[key] = value
		return value
	}
}
