package main

import (
	"fmt"
	"os"
	"unicode/utf8"
)

// IsPalindrome ...
var IsPalindrome func(string) bool // Holds a reference to a function

func init() {
	if len(os.Args) > 1 && (os.Args[1] == "-a" || os.Args[1] == "--ascii") {
		os.Args = append(os.Args[:1], os.Args[2:]...) // stripg out arg
		IsPalindrome = func(s string) bool {          // Simple ASCII-only version
			if len(s) <= 1 {
				return true
			}
			if s[0] != s[len(s)-1] {
				return false
			}
			return IsPalindrome(s[1 : len(s)-1])
		}
	} else {
		IsPalindrome = func(s string) bool { // UTF-8 Version
			if utf8.RuneCountInString(s) <= 1 {
				return true
			}

			first, sizeOfFirst := utf8.DecodeRuneInString(s)
			last, sizeOfLast := utf8.DecodeLastRuneInString(s)
			if first != last {
				return false
			}

			return IsPalindrome(s[sizeOfFirst : len(s)-sizeOfLast])
		}
	}
}

func main() {
	words := []string{"PULLUP", "ROTOR", "DECIDED"}
	for _, word := range words {
		fmt.Printf("%s is palindrome: %t\n", word, IsPalindrome(word))
	}
}
